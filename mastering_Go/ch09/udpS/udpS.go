package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("plz provide a port number")
		return
	}
	PORT := ":" + arguments[1]
	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	buffer := make([]byte, 1024)

	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s --> %s\n", addr, string(buffer[:n]))

		if strings.TrimSpace(string(buffer[:n])) == "STOP" {
			fmt.Println("exit server")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data : %s\n", string(data))

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

}
