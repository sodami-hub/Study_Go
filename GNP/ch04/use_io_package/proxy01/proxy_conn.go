/*
io 패키지에서 가장 유용한 함수 중 하나는 io.Reader에서 데이터를 읽어서 io.Writer로 데이터를 쓸 수 있는 io.Copy 함수이다.
이 함수는 두 노드 중간에서 데이터를 전송하는 프락시(proxy)를 생성하는데 유용하다.

아래 코드의 proxyConn 함수는 두 노드 간의 연결로부터 프락시를 생성하는 것은 아주 쉽다. 이 함수는 출발지 노드로부터 전송된 데이터를 목적지 노드로,
목적지 노드에서 전송된 데이터를 출발지 노드로 복제한다.

두 노드가 서로 직접 연결한 것처럼 프락시로 데이터를 주고받을 수 있다.
*/

package main

import (
	"io"
	"net"
)

func proxyConn(source, destination string) error { // source -> Writer(출발지) / destination -> Reader(목적지)
	connSource, err := net.Dial("tcp", source)
	if err != nil {
		return err
	}
	defer connSource.Close()

	connDestination, err := net.Dial("tcp", destination)
	if err != nil {
		return err
	}
	defer connDestination.Close()

	// connSource에 대응하는 connDestination
	go func() { _, _ = io.Copy(connSource, connDestination) }()

	// connDestination으로 메시지를 보내는 connSource
	_, err = io.Copy(connDestination, connSource)

	return err
}
