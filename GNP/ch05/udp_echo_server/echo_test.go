package echo

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"testing"
)

func echoServerUDP(ctx context.Context, addr string) (net.Addr, error) {

	// UDP  연결에는 Accept 이런거 없다. 그냥 listner 열어두면 여기로 그냥 다 보낸다..?
	s, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = s.Close()
		}()

		buf := make([]byte, 1024)

		for {
			n, clientAddr, err := s.ReadFrom(buf) // 클라이언트에서 서버로
			if err != nil {
				return
			}

			_, err = s.WriteTo(buf[:n], clientAddr) // 서버에서 클라이언트로 에코잉
			if err != nil {
				return
			}
		}
	}()
	return s.LocalAddr(), nil
}

func TestEchoSErverUDP(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	msg := []byte("ping")
	_, err = client.WriteTo(msg, serverAddr)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, addr, err := client.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	if addr.String() != serverAddr.String() {
		t.Fatalf("received reply form %q instead of %q", addr, serverAddr)
	}

	if !bytes.Equal(msg, buf[:n]) {
		t.Errorf("expected reply %q; actual reply %q", msg, buf[:n])
	}
}
