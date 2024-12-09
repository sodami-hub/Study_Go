/*
여러개의 클라이언트가 연결될 수 있는 tcp 서버
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	var wg sync.WaitGroup

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			break
		}

		wg.Add(1)
		go func(c net.Conn) {
			defer wg.Done()
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
				fmt.Println(err)
				return
			}

			// 고루틴과 클라이언트 접속 종료... 어떻게 처리?
			if strings.TrimSpace(string(netData)) == "STOP" {
				fmt.Println("Exiting Client!")
				return
			}

			fmt.Println("-> ", string(netData))
			t := time.Now()
			myTime := t.Format(time.RFC3339) + "\n"
			c.Write([]byte(myTime))
		}(c)
	}
	wg.Wait()
	fmt.Println("server Exiting...")
}
