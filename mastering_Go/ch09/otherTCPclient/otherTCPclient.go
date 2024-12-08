package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("provide a server:port")
		return
	}

	connect := arguments[1]

	// net.ResolveTCPAddr() 함수는 TCP 네트워크에서만 사용할 수 있고 주어진 주소에 대한 *net.TCPAddr 타입의 값을 반환한다.
	// *net.TCPAddr 구조체는 TCP 엔드포인트의 주소를 나타낸다.
	// net.ResolveTCPAddr()과 net.DialTCP를 사용하는 방법은 net.Dial()을 사용하는 것보다 복잡하지만 더 많은 제어를 할 수 있게 해준다.
	tcpAddr, err := net.ResolveTCPAddr("tcp4", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	// TCP 앤드포인트를 알고 있으니 이를 활용해 net.DialTCP()로 서버에연결한다.
	conn, err := net.DialTCP("tcp4", nil, tcpAddr)

	// conn, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, "from Client : "+text+"\n")

		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("from Server : " + message)
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}
}
