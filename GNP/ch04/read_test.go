// 고정된 버퍼에 데이터 읽기

package main

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

func TestReadIntoBuffer(t *testing.T) {
	payload := make([]byte, 1<<24) // 16MB 페이로드 생성
	_, err := rand.Read(payload)   // radn.Read(payload) []byte 에 의사난수를 채움
	if err != nil {
		t.Fatal(err)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload)
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1<<19) // 512KB 버퍼
	count := 1
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Error(err)
			}
			break
		}
		t.Logf("%d : read %d bytes", count, n) // n은 읽어드린 버퍼의 크기 / 버퍼는 바이트 슬라이스 이기 때문에 바이트의 크기이다.
		count++
	}
	conn.Close()
}
