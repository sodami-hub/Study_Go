// TCP 연결 시도 시 타임아웃 기간 설정하기

package ch03

import (
	"net"
	"syscall"
	"testing"
	"time"
)

func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error { // net.Dialer구조체의 Control필드(메서드필드)에 대한 정의 구조체를 초기화할 때 사용 examStruct := struct{age:15,...}
			return &net.DNSError{
				Err:         "connecton time out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout, // net.Dialer의 Timeout에 대한 정의
	}
	return d.Dial(network, address)
}

func TestDialTimeout(t *testing.T) {
	c, err := DialTimeout("tcp", "10.0.0.1:http", 5*time.Second)
	if err == nil {
		c.Close()
		t.Fatal("connection did not time out")
	}
	nErr, ok := err.(net.Error) // 에러가 발생했을 때
	if !ok {
		t.Fatal(err)
	}
	if !nErr.Timeout() { //
		t.Fatal("error is not timeout")
	}
}
