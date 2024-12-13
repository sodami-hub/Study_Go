package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/websocket"
)

var PORT = ":1234"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!\n")
	fmt.Fprintf(w, "Please use /ws for WebSocket")
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("connection from", r.Host)

	// 웹소켓 서버 어플리케이션은 Upgrader.Upgrade 메서드를 호출해 HTTP 요청에서 웹소켓 연결로 업그레이드한다.
	// Upgrader.Upgrade 요청이 성공하면 서버는 웹소켓 연결을 사용해 클라이언트와 통신을 시작한다.
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrader.Upgrade: ", err)
		return
	}
	defer ws.Close()

	// 아래 for 루프를 통해 /ws로 들어오는 모든 메세지를 처리한다.
	// 중요한 것은 웹소켓 연결에서 클라이언트로 데이터를 보낼 때 fmt.Fprintf 를 사용할 수 없다는 것을 기억해야 된다.
	// gorilla/websocket 을 사용할때는 WriteMessage()와 ReadMessage()만을 사용해야 한다.
	for {
		mt, message, err := ws.ReadMessage()
		if err != nil {
			log.Println("FROM", r.Host, "read", err)
			break
		}
		log.Print("Received: ", string(message))
		err = ws.WriteMessage(mt, message)
		if err != nil {
			log.Println("WriteMessage : ", err)
			break
		}
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/", http.HandlerFunc(rootHandler))
	mux.Handle("/ws", http.HandlerFunc(wsHandler))

	log.Println("Listening to TCP Port", PORT)
	err := s.ListenAndServe()
	if err != nil {
		log.Println(err)
		return
	}
}
