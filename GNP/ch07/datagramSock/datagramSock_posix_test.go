//go:build darwin || linux

/*
unixgram은 Window에서 동작하지 않기 때문에 build constraint를 이용하여 이 코드가 Windows에서 동작하지 않도록 하고
패키지를 플랫폼에 맞게 올바르게 임포트할 수 있도록 한다.

데이터 그램을 처리하기 때문에 서버가 응답한 ping은 개별 메시지의 경계가 명확하고 따라서 마지막에 읽은 값은 ping이다.
*/

package datagramsock

import (
	"bytes"
	"context"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"testing"
)

func TestEchoServerUnixDatagram(t *testing.T) {

	// ================================== 데이터그램 기반 에코 서버 초기화
	dir, err := os.MkdirTemp("", "echo_unixgram")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if rErr := os.RemoveAll(dir); rErr != nil {
			t.Error(rErr)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	sSocket := filepath.Join(dir, fmt.Sprintf("s%d.sock", os.Getpid()))
	serverAddr, err := datagramEchoServer(ctx, "unixgram", sSocket)
	if err != nil {
		t.Fatal(err)
	}
	defer cancel()

	err = os.Chmod(sSocket, os.ModeSocket|0622)
	if err != nil {
		t.Fatal(err)
	}

	// ==================================== 데이터그램 기반 클라이언트 초기화
	cSocket := filepath.Join(dir, fmt.Sprintf("c%d.sock", os.Getpid()))
	client, err := net.ListenPacket("unixgram", cSocket)
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = client.Close() }()

	err = os.Chmod(cSocket, os.ModeSocket|0622)
	if err != nil {
		t.Fatal(err)
	}

	// ==================================== unixgram 소켓을 이용한 메시지 에코잉
	msg := []byte("ping")
	for i := 0; i < 3; i++ {
		_, err := client.WriteTo(msg, serverAddr)
		if err != nil {
			t.Fatal(err)
		}
	}

	buf := make([]byte, 1024)
	n, addr, err := client.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	if addr.String() != serverAddr.String() {
		t.Fatalf("received reply from %q instead of %q", addr, serverAddr)
	}

	t.Log("from server :", string(buf[:n]))

	if !bytes.Equal(msg, buf[:n]) {
		t.Fatalf("expected reply %q; actual reply %q", msg, buf[:n])
	}
}
