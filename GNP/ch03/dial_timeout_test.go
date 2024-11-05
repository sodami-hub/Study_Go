// TCP 연결 시도 시 타임아웃 기간 설정하기

/*
net.Dialer의 Control 필드는 네트워크 연결을 설정할 때 항상 호출되는 함수입니다. 이 함수는 네트워크 연결을 제어하는 데 사용됩니다. Control 필드를 설정하면, Dial 또는 DialContext 메서드가 호출될 때마다 이 함수가 호출됩니다.

Control 필드의 역할
Control 필드는 네트워크 연결을 설정할 때 호출되는 함수로, 연결 설정을 제어할 수 있습니다.
*/
package ch03

import (
	"net"
	"syscall"
	"testing"
	"time"
)

/*
net.Dialer 구조체를 초기화합니다.
Control 필드는 네트워크 연결을 제어하는 함수입니다. 여기서는 항상 타임아웃 에러를 반환하도록 설정되어 있습니다.
net.DNSError 구조체를 반환하여 타임아웃 에러를 시뮬레이션합니다.
Err: 에러 메시지입니다.
Name: 주소 이름입니다.
Server: 서버 주소입니다.
IsTimeout: 타임아웃 여부를 나타내는 불리언 값입니다.
IsTemporary: 임시 에러 여부를 나타내는 불리언 값입니다.
Timeout 필드는 연결 시도 시 적용할 타임아웃 시간을 설정합니다.
*/
func DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error { // net.Dialer구조체의 Control필드(메서드필드)에 대한 정의 구조체를 초기화할 때 사용 examStruct := struct{age:15,...}
			return &net.DNSError{ // 연결 시도시 항상 DNSError을 반환함
				Err:         "connecton time out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout, // net.Dialer의 Timeout에 대한 정의
	}
	return d.Dial(network, address) // 연결 시도.
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

/*
DialTimeout - 함수는 네트워크 연결을 설정할 때 타임아웃을 적용하는 함수입니다.
이 함수는 net.Dialer 구조체를 사용하여 네트워크 연결을 설정하고, 타임아웃이 발생하면 에러를 반환합니다.
각 부분의 역할을 설명하겠습니다.

이 함수는 네트워크 연결을 설정할 때 타임아웃을 적용하는 방법을 보여줍니다.
net.Dialer 구조체를 사용하여 타임아웃을 설정하고, 타임아웃이 발생하면 에러를 반환합니다.
이 함수는 주어진 네트워크와 주소로 연결을 시도하며, 타임아웃이 발생하면 타임아웃 에러를 반환합니다.

*/
