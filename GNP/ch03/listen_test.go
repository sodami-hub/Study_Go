package main

import (
	"net"
	"testing"
)

func TestLinstener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0") // net.Listen 함수는 네트워크의 종류와 IP주소:포트 문자열을 매개변수로 받는다.
	// 반환값으로 net.Listener 인터페이스와 에러 인터페이스를 반환받는다.
	// IP주소:포트 에 대한 인수는 비워둘 수 있다. Go가 무작위 포트 번호를 할당한다. Addr 메서드를 이용해서리스너의 주소를 얻어 올 수 있다.
	// 포트 번호를 0으로 할당하면 무작위 할당
	// "" 이렇게 하면 시스템상의 모든 유니캐스트와 애니캐스트 IP주소에 바인딩된다.
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = listener.Close() }() // 리스너 우아하게 종료하기.

	t.Logf("bound to %q", listener.Addr()) // 리스너 주소 얻어오기
}
