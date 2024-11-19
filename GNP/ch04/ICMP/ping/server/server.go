/*
ping.go 를 테스트 하기 위한 서버 리스너
*/

package main

import (
	"fmt"
	"net"
	"sync"
)

var wg sync.WaitGroup

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("접속 대기 에러")
		return
	}
	wg.Add(1)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			defer conn.Close()
			fmt.Println("서버-클라이언트")
			buf := make([]byte, 1024)
			for {
				_, err := conn.Read(buf)
				if err != nil {
					fmt.Println("읽기 에러:", err)
					wg.Done()
					return
				}
			}
		}
	}()

	wg.Wait()
}
