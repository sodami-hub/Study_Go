/*
네트워크 연결 모니터링
io 패키지를 이용하면 연결 객체를 통해 네트워크 데이터를 주고받는 것 이상으로 유용한 동작을 할 수 있다.
예를 들어, io.MultiWriter 함수를 이용하여 단일 페이로드를 여러 개의 네트워크 연결로 전송할 수 있다.
io.TeeReader 함수를 사용하여 네트워크 연결로부터 읽은 데이터를 로깅하는 데 사용할 수도 있다.

아래 코드는 io.TeeReader와 io.MultiWriter 함수를 사용하여 TCP리스너로부터 발생하는 모든 네트워크 트래픽을 로깅하는 예시를 보여준다.
*/

package main

import (
	"io"
	"log"
	"net"
	"os"
)

// Monitor 구조체는 네트워크 트래픽을 로깅하기 위한 log.Logger를 임베딩한다.
type Monitor struct {
	*log.Logger
}

// Writer 메서드는 io.Writer 인터페이스를 구현한다.
func (m *Monitor) Write(p []byte) (int, error) {
	return len(p), m.Output(2, string(p))
}

func main() {
	// 표준 출력으로 데이터를 쓰는 Monitor 구조체의 인스턴스를 만든다.
	monitor := &Monitor{Logger: log.New(os.Stdout, "monitor: ", 0)}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		monitor.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		defer close(done)

		conn, err := listener.Accept()
		if err != nil {
			return
		}
		defer conn.Close()

		b := make([]byte, 1024)
		r := io.TeeReader(conn, monitor) // TeeReader 에 Reader는 conn, Writer는 monitor 전달
		n, err := r.Read(b)              // 네트워크 연결에서 데이터를 읽고, 읽은 데이터를 모니터에 출력
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}

		w := io.MultiWriter(conn, monitor) //네트워크 연결에 데이터를 쓰고, 쓴 데이터를 모니터에 출력

		_, err = w.Write(b[:n]) // 메시지를 에코잉한다.
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		monitor.Fatal(err)
	}

	// 서버에서 데이터를 읽을 때, 다시 클라이언트로 보낼때 -> 화면에 두번 출력될 것이다.
	_, err = conn.Write([]byte("hello world\n"))
	if err != nil {
		monitor.Fatal(err)
	}
	_ = conn.Close()

	<-done
}
