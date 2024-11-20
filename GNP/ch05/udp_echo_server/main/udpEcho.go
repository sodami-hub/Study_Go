package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

func echoServerUDP(ctx context.Context, addr string) (net.Addr, error) {

	// UDP  연결에는 Accept 이런거 없다. 그냥 listner 열어두면 여기로 그냥 다 보낸다..?
	s, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}

	go func() {
		go func() {
			fmt.Println("cancel 메세지 수신 대기중")
			<-ctx.Done()
			fmt.Println("cancel 메세지 수신")
			_ = s.Close()
		}()

		buf := make([]byte, 1024)

		for {
			fmt.Println("서버 읽기 블로킹 들어감")
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	serverAddr, err := echoServerUDP(ctx, "127.0.0.1:")
	if err != nil {
		log.Printf("%v", err)
		return
	}
	// context의 취소함수 -> 서버 또한 종료가 됨.
	defer cancel()

	// 클라이언트의 연결객체를 생성하고  WriteTo() 를 호출할 때마다. 매개변수로 주소를 전달해야 됨.
	client, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	defer func() { _ = client.Close() }()

	msg := []byte("ping ping ping")

	fmt.Println("클라이언트 슬립 3초")
	time.Sleep(3 * time.Second)

	_, err = client.WriteTo(msg, serverAddr)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Printf("send server : %s\n", msg)

	buf := make([]byte, 1024)
	n, addr, err := client.ReadFrom(buf)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}
	fmt.Printf("receive server : %s\n", buf[:n])

	if addr.String() != serverAddr.String() {
		log.Printf("received reply form %q instead of %q", addr, serverAddr)
		return
	}

	if !bytes.Equal(msg, buf[:n]) {
		log.Printf("expected reply %q; actual reply %q", msg, buf[:n])
		return
	}
}
