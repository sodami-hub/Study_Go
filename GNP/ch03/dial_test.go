// 127.0.0.1로 연결 수립하기

package ch03

import (
	"io"
	"net"
	"testing"
)

func TestDial(t *testing.T) {
	// 랜덤 포트에 리스너 생성
	listener, err := net.Listen("tcp", "127.0.0.1:0") // 리스너 생성
	if err != nil {
		t.Fatal("create listener err", err)
	}

	done := make(chan struct{}) // 다른고루틴과 통신할 채널 생성(빈 구조체를 주고 받음)

	go func() {
		defer func() { done <- struct{}{} }() // 고루틴을 닫을 때 빈구조체 인스턴스를 채널에 전송

		for {
			conn, err := listener.Accept() // 고루틴에서 클라이언트로부터의 연결을 대기
			if err != nil {
				t.Log(err)
				return
			}

			go func(c net.Conn) {
				defer func() {
					c.Close()
					done <- struct{}{}
				}()

				buf := make([]byte, 1024)
				for {
					n, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
					t.Logf("received: %q", buf[:n])
				}
			}(conn)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	conn.Close()
	<-done // 다른 고루틴에서 done 채널로 값을 보낼때 까지 대기 - 보낸 값에 대해서는 무시.
	listener.Close()
	<-done
}
