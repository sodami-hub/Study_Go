/*
net.Conn 인터페이스를 구현하는 UDP 객체를 사용하여 연결을 수립할 수 있다. 그러면 TCP에서 구현하는 코드와 동일한 코드를 사용할 수 있다.
UDP 기반 연결로 net.Conn 인터페이스를 사용하면 끼어든 연결에서 인터럽트 메시지를 받을 필요도 없으며, 따라서 응답마다 송신자의 주소를 확인할 필요도 없다.
*/

package echo

import (
	"bytes"
	"context"
	"net"
	"testing"
	"time"
)

func TestDialUDP(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:") // 서버 udp리스너의 주소를 가져온다. net.PacketConn
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	client, err := net.Dial("udp", serverAddr.String()) // net.Conn을 사용해서 서버 net.PacketConn에 연결
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	interloper, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	// 서버의 응답전에 클라이언트에 메세지를 보냄(클라이언트 인터럽트)
	interrupt := []byte("pardon me")
	n, err := interloper.WriteTo(interrupt, client.LocalAddr()) // 클라이어트로 메시지가 들어가지 않는다. 현재 클라이언트는 서버와만 연결돼있다.(서버측은 아님...)
	if err != nil {
		t.Fatal(err)
	}
	defer interloper.Close()

	if l := len(interrupt); l != n {
		t.Fatalf("wrote %d bytes of %d", n, l)
	}
	// ========

	// net.PacketConn을 이용한 UDP 연결과 net.Conn을 이용한 UDP 연결의 차이
	ping := []byte("ping")
	_, err = client.Write(ping)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, err = client.Read(buf) // 읽어올때 송신자의 주소를 가져오지 않는다.
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(ping, buf[:n]) {
		t.Errorf("expected reply %q; actual reply %q", ping, buf[:n])
	}

	err = client.SetDeadline(time.Now().Add(3 * time.Second))
	if err != nil {
		t.Fatal(err)
	}

	_, err = client.Read(buf) // 데드라인(3초) 동안 기다려본다.
	if err == nil {
		t.Fatal("unexpected packet")
	}
}
