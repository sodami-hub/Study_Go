package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var port int
	var addr string
	var nick string

	// 1.실행 인수를 통해 접근하려는 ip, port 설정
	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.StringVar(&addr, "address", "localhost", "--address localhost")

	// 2 net 패키지를 이용해서 tcp 연결을 맺는다.
	tcpAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		log.Fatal("ResolveTCPAddr failed:", err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal("Dial failed:", err)
	}

	fmt.Print("닉네임을 입력하세요 : ") // 닉네임을 입력받음
	_, err = fmt.Scanf("%s\n", &nick)
	if err != nil {
		log.Fatal("make nickname error")
	}

	// 3. 연결된 conn을 통해 데이터를 읽어서 출력한다.
	go func() {
		scan := bufio.NewScanner(conn)
		scan.Split(bufio.ScanLines)
		for scan.Scan() {
			fmt.Println(scan.Text())
		}
	}()

	// 4. 키보드로부터 텍스트를 입력받아 데이터를 전송한다.
	for {
		inputScan := bufio.NewScanner(os.Stdin)
		inputScan.Split(bufio.ScanLines)
		for inputScan.Scan() {
			if inputScan.Text() == "exit" {
				return
			}
			conn.Write([]byte(fmt.Sprintf("%s : %s\n", nick, inputScan.Text()))) // 텍스트 전송할 때 닉네임을 붙여서 전송
		}
	}
}
