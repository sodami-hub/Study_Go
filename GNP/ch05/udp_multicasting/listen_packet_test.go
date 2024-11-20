package echo

import (
	"bytes"
	"context"
	"net"
	"testing"
)

func TestListenPacketUDP(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:") // 에코 서버 생성
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	client, err := net.ListenPacket("udp", "127.0.0.1:") // 클라이언트 생성
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	interloper, err := net.ListenPacket("udp", "127.0.0.1:") // 클라이언트와 에코 서버 사이에 끼어들 새로운 UDP 연결 생성
	if err != nil {
		t.Fatal(err)
	}

	interrupt := []byte("pardon me")
	n, err := interloper.WriteTo(interrupt, client.LocalAddr()) // 클라이언트를 인터럽트.. 이 메세지는 클라이언트의 수신 버퍼에 큐잉
	if err != nil {
		t.Fatal(err)
	}
	_ = interloper.Close()

	if l := len(interrupt); l != n {
		t.Fatalf("wrote %d bytest of %d", n, l)
	}

	ping := []byte("ping")
	_, err = client.WriteTo(ping, serverAddr)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, addr, err := client.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(interrupt, buf[:n]) {
		t.Errorf("expected reply %q; actual reply %q", interrupt, buf[:n])
	}
	t.Logf("reply : %q", buf[:n])

	if addr.String() != interloper.LocalAddr().String() {
		t.Errorf("expected message from %q; actual sender is %q", interloper.LocalAddr(), addr)
	}

	n, addr, err = client.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(ping, buf[:n]) {
		t.Errorf("expected reply %q; actual reply %q", ping, buf[:n])
	}

	if addr.String() != serverAddr.String() {
		t.Errorf("expected message from %q; actual sender is %q", serverAddr, addr)
	}
}
