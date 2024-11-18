/*
proxy01의 proxyConn 함수를 사용하는 약간 다른 방법. proxy01은 두 노드 간에 네트워크 연결을 수립하고 트래픽을 프락시하는 반면

아래 코드는 io.Reader와 io.Writer 인터페이스 간에 데이터를 프락시하여, 네트워크 연결 외의 것들에도 적용할 수 있기 때문에 테스트하기에도 훨씬 쉽다.
*/

package main

import (
	"io"
	"net"
	"sync"
	"testing"
)

// net.Conn 인터페이스 대신 범용적인 io.Reader/Writer 인터페이스를 매개변수로 받기 때문에 조금 더 활용 범위가 넓다.
// 이를 사용하여 데이터를 네트워크 연결로부터 os.Stdout, *bytes.Buffer. *os.File 외에 io.Writer 인터페이스를 구현한 많은 객체들로 데이터를 프락시할 수 있다.
func proxy(from io.Reader, to io.Writer) error {
	fromWriter, fromIsWriter := from.(io.Writer)
	toReader, toIsReader := to.(io.Reader)

	if toIsReader && fromIsWriter {
		// 필요한 인터페이스를 모두 구현하였으니 from과 to에 대응하는 프락시 생성
		go func() { _, _ = io.Copy(fromWriter, toReader) }()
	}

	_, err := io.Copy(to, from)

	return err
}

// proxy 함수가 의도대로 동작하는지 확인하는 테스트 생성
func TestProxy(t *testing.T) {
	var wg sync.WaitGroup

	//서버는 핑 메시지를 대기하고 퐁 메시지로 응답
	// 그 외의 메시지는 동일하게 클라이언트로 에코잉된다.
	server, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			// 클라이언트 접속 대기
			conn, err := server.Accept()
			if err != nil {
				return
			}
			// 클라이언트 접속 후 서버측 처리
			go func(c net.Conn) {
				defer c.Close()

				for {
					buf := make([]byte, 1024)
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					// ping을 받으면 pong을 보내고, 아니면 그대로 다시 보내주는 에코 서버
					switch msg := string(buf[:n]); msg {
					case "ping":
						_, err = c.Write([]byte("pong"))
					default:
						_, err = c.Write(buf[:n])
					}

					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
				}
			}(conn)
		}
	}()

	// proxyServer는 메시지를 클라이언트 연결로부터 destinationServer로 프락시한다.
	// destinationServer 서버에서 온 응답 메시지는 역으로 클라이언트에 프락시된다.
	proxyServer, err := net.Listen("tcp", "127.0.0.1:") // 프록시 서버 세팅
	if err != nil {
		t.Fatal(err)
	}

	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			conn, err := proxyServer.Accept() // 클라이언트 <-> 프록시 => from
			if err != nil {
				return
			}
			go func(from net.Conn) {
				defer from.Close()

				// 프록시 <-> 서버 => to
				to, err := net.Dial("tcp", server.Addr().String())
				if err != nil {
					t.Error(err)
					return
				}

				defer to.Close()

				err = proxy(from, to)
				if err != nil && err != io.EOF {
					t.Error(err)
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", proxyServer.Addr().String()) // 클라이언트의 프록시서버 연결 요청 -> line 99
	if err != nil {
		t.Fatal(err)
	}

	msgs := []struct{ Message, Reply string }{
		{"ping", "pong"},
		{"ping", "pong"},
		{"ping", "pong"},
		{"ping", "pong"},
	}

	for i, m := range msgs {
		_, err := conn.Write([]byte(m.Message))
		if err != nil {
			t.Fatal(err)
		}

		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(buf[:n])
		t.Logf("%q -> proxy -> %q", m.Message, actual)

		if actual != m.Reply {
			t.Errorf("%d: expected reply: %q; actual: %q", i, m.Reply, actual)
		}
	}

	_ = conn.Close()
	_ = proxyServer.Close()
	_ = server.Close()

	wg.Wait()

}
