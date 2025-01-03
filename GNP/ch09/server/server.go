package main

import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"GNP/handlers"
)

var wg sync.WaitGroup

func main() {
	srv := &http.Server{
		Addr: "127.0.0.1:8081",
		Handler: http.TimeoutHandler(
			handlers.DefaultHandler(), 2*time.Minute, ""),
		IdleTimeout:       5 * time.Minute,
		ReadHeaderTimeout: time.Minute,
	}

	l, err := net.Listen("tcp", srv.Addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		/*
		   Serve() : http.Server 의 메서드는 서버가 종료될 때까지 블록처리된다.(고수준 서버 프로그래밍에서 사용)
		   이 함수는 HTTP 서버를 시작하고 클라이언트의 요청을 처리하는 역할을 한다.
		   net.Listener를 인수로 받아들이며 내부적으로 Accept() 함수를 호출하고, HTTP 요청을 처리한다.

		   Accept() : net.Listener 인터페이스의 메서드로, 새로운 클라이언트 연결을 수락하는 역할을 한다.(저수준 서버 프로그래밍에서 사용)
		   이 함수는 블로킹 호출로, 새로운 연결이 들어올 때까지 대기한다. 새로운 연결이 들어오면,
		   net.Conn 인터페이스를 반환하여 클라이언트와의 통신을 처리할 수 있게 한다.
		*/
		err := srv.Serve(l)
		if err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()
	wg.Wait()
}
