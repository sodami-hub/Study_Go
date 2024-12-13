/*
-> 먼저 /ws/ws.go => 웹소켓 서버를 실행하고
go run wsClient.go localhost:1234 ws => 클라이언트를 실행한다.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var SERVER = ""
var PATH = ""
var TIMESWAIT = 0
var TIMESWAITMAX = 5
var in = bufio.NewReader(os.Stdin)

// getInput() 함수는 고루틴 형태로 실행하고 사용자 입력을 받아 input 채널을 통해 main()함수로 전송한다.
// 프로그램이 사용자 입력을 받을 때마다 예전 고루틴은 종료되고 새로운 getInput() 고루틴이 시작돼 사용자 입력을 받는다.
func getInput(input chan string) {
	result, err := in.ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	input <- result
}

func main() {
	arguments := os.Args
	if len(arguments) != 3 {
		fmt.Println("Need SERVER + PATH!")
		return
	}

	SERVER = arguments[1]
	PATH = arguments[2]
	fmt.Println("Connecting to :", SERVER, "at", PATH)

	// 웹소켓 클라이언트는 interrupt 채널을 이용해 유닉스 인터럽트를 처리한다.
	// 적절한 시그널이 들어오면 websocket.CloseMessage를 사용해 서버와의 웹소켓 연결이 종료된다.
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	input := make(chan string, 1)
	go getInput(input)

	// 웹소켓 연결은 websocket.DefaultDialer.Dial()을 호출해 시작한다. input 채널로 입력하는 모든 메시지는 WriteMessage() 메서드를 통해 웹소켓 서버로 전송한다.
	URL := url.URL{Scheme: "ws", Host: SERVER, Path: PATH}
	c, _, err := websocket.DefaultDialer.Dial(URL.String(), nil)
	if err != nil {
		log.Println("Error:", err)
		return
	}
	defer c.Close()

	done := make(chan struct{})
	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("ReadMessage error: ", err)
				return
			}
			log.Printf("Received: %s", message)
		}
	}()

	for {
		select {
		case <-time.After(4 * time.Second):
			log.Println("plz give me input", TIMESWAIT)
			TIMESWAIT++
			if TIMESWAIT > TIMESWAITMAX {
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			}
		case <-done:
			return
		case t := <-input:
			err := c.WriteMessage(websocket.TextMessage, []byte(t))
			if err != nil {
				log.Println("Write error:", err)
				return
			}
			TIMESWAIT = 0
			go getInput(input)
		case <-interrupt:
			log.Println("Caought interrupt signal - quitting!")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("Write close error :", err)
				return
			}
			select {
			case <-done:
			case <-time.After(2 * time.Second):
			}
			return
		}
	}

}
