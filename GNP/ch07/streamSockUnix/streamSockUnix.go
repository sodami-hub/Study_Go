/*
스트리밍 소켓 -unix
스트리밍 유닉스 도메인 소켓은 TCP가 갖는 메시지 확인, 체크섬, 혼잡 제어 등의 오버헤드 없이도 TCP 처럼 동작한다.
운영체제는 책임지고 유닉스 도메인 소켓을 이용하여 TCP처럼 동작하는 IPC를 구현한다.


*/

package streamsockunix

import (
	"context"
	"net"
)

func streamingEchoServer(ctx context.Context, network string, addr string) (net.Addr, error) {
	// network 타입에는 tcp, unix, unixpacket 같은 문자열을 전달 받는다.
	// addr은 tcp 인경우는 IP주소와 포트번호의 조합으로, unix, unixpacket의 경우는 존재하지 않는 파일의 경로여야 한다.
	// 에코 서버가 바인딩을 하게 되면 소켓 파일이 생성된다. 이후 서버는 연결 요청을 수신 대기한다.
	s, err := net.Listen(network, addr)
	if err != nil {
		return nil, err
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = s.Close()
		}()

		for {
			conn, err := s.Accept()
			if err != nil {
				return
			}

			go func() {
				defer func() { _ = conn.Close() }()

				for {
					buf := make([]byte, 1024)
					n, err := conn.Read(buf)
					if err != nil {
						return
					}

					_, err = conn.Write(buf[:n])
					if err != nil {
						return
					}
				}
			}()
		}
	}()

	return s.Addr(), nil
}
