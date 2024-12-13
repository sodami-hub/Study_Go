package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("need socket path")
		return
	}
	socketPath := os.Args[1]

	c, err := net.Dial("unix", socketPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		_, err := c.Write([]byte(text))
		if err != nil {
			fmt.Println("Write:", err)
			break
		}
		buf := make([]byte, 256)

		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Read : ", buf[:n])

		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("exiting unix domain socket client!")
			return
		}

		time.Sleep(5 * time.Second)
	}
}
