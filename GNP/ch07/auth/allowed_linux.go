package auth

import (
	"log"
	"net"
	"os/user"
	"strconv"

	"golang.org/x/sys/unix"
)

func Allowed(conn *net.UnixConn, groups map[string]struct{}) bool {
	if conn == nil || groups == nil || len(groups) == 0 {
		return false
	}
	// net.UnixConn : Go에서의 TCP 연결 객체를 나타내는 net.TCPConn 객체와도 유사하다. 연결 객체로부터 파일 디스크립터 정보를 획득해야 되므로 리스너의 Accept() 메서드로부터 반환된
	// net.Conn 인터페이스만드로는 충분하지 않다. 대신에 리스너의 AcceptUnix 메서드로부터 반환된 net.UnixConn 객체의 포인터를 이 함수로 넘겨준다.
	file, _ := conn.File() // 피어의 유닉스 인증 정보를 얻어 오기 위해 먼저 net.UnixConn 파일 객체를 변수로 저장한다.
	defer func() { _ = file.Close() }()

	var (
		err   error
		ucred *unix.Ucred
	)

	for {
		// unix.GetsockoptUcred 함수의 매개변수로 파일 객체의 디스크립터, 어느 프로토콜 계층에 속했는지 나타내는 상수인 unix.SOL_SOCKET, 그리고 옵션값(unix.SO_PEERCRED) 를 넘겨준다.
		// 리눅스 커널에서 소켓 옵션값을 얻어 오려면 해당하는 옵션과 해당 옵션이 존재하는 계층 값이 모두 필요하다. unix.SOL_SOCKET 값은 리눅스 커널에 소켓 계층의 옵션 값이 필요하다고 알려주며,
		// 마찬가지로 unix.SOL_TCP 값은 리눅스 커널에 TCP 계층의 옵션 값이 필요하다고 알려준다. unix.SO_PEERCRED 상숫값은 리눅스 커널에 피어의 인증 정보가 필요하다고 알려준다.
		// 리눅스 커널이 유닉스 도메인 소켓 계층의 피어 인증 정보를 찾으면 unix.GetsockoptUcred 함수는 정상적인 unix.Ucred 객체의 포인터를 반환한다.
		ucred, err = unix.GetsockoptUcred(int(file.Fd()), unix.SOL_SOCKET, unix.SO_PEERCRED)
		if err == unix.EINTR {
			continue // syscall 중단됨, 다시 시도하기
		}
		if err != nil {
			log.Println(err)
			return false
		}
		break
	}
	// unix.Ucred 객체에는 피어의 프로세스 정보와 사용자 ID, 그룹 ID 정보가 있다. 피어의 사용자 ID를(uID) 매개변수로 전달한다.
	u, err := user.LookupId(strconv.Itoa(int(ucred.Uid)))
	if err != nil {
		log.Println(err)
		return false
	}

	// user.LookupID 를 통해서 반환받은 사용자 정보를 통해서 그룹 ID의 목록을 반환한다. 사용자는 하나 이상의 그룹에 속할 수 있다.
	gids, err := u.GroupIds()
	if err != nil {
		log.Println(err)
		return false
	}

	// 사용자가 속한 그룹 아이디와 접속이 허용된 그룹 아이디 groups 를 비교한다. 사용자의 그룹 아이디가 groups에 속해있으면 true를 반환한다.
	for _, gid := range gids {
		if _, ok := groups[gid]; ok {
			return true
		}
	}
	return false
}
