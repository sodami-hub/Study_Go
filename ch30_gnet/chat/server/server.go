// echo 서버를 기반으로 chatserver 만들기 . ehco 서버는 수신한 데이터를 수신자에게만 전송하지만 채팅 서버는 연결된 모든 클라이언트에게 전송한다.(브로트캐스트)
package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	gnet "github.com/panjf2000/gnet/v2"
)

type chatServer struct {
	gnet.BuiltinEventEngine

	// 1. 연결된 커넥션을 보관하는 맵
	cliMap sync.Map
}

func (cs *chatServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	log.Printf("client connected. address:%s", c.RemoteAddr().String())

	// 2. 새로운 커넥션 보관
	cs.cliMap.Store(c, true)
	return nil, gnet.None
}

func (cs *chatServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	log.Printf("cliend disconnected. address:%s", c.RemoteAddr().String())

	// 3. 연결 해제된 클라이언트 맵에서 삭제
	if _, ok := cs.cliMap.LoadAndDelete(c); ok {
		log.Printf("connection removed")
	}
	return gnet.None
}

func (cs *chatServer) OnBoot(eng gnet.Engine) gnet.Action {
	log.Printf("chat server is listening on\n")
	return gnet.None
}

func (cs *chatServer) OnTraffic(c gnet.Conn) gnet.Action {
	buf, _ := c.Next(-1) // 모든 데이터를 읽는다.
	// 4. 데이터 수신 시 모든 커넥션에 데이터 전송
	cs.cliMap.Range(func(key, value any) bool {
		if conn, ok := key.(gnet.Conn); ok {
			if conn != c { // 나 이외의 다른 클라이언트에게만 전송되도록 수정
				msg := []byte(conn.LocalAddr().String()) // 메시지 앞에 ip 주소 붙여봄
				msg = append(msg, buf...)
				conn.AsyncWrite(msg, nil)
			}
		}
		return true
	})
	return gnet.None
}

func main() {
	var port int
	var multicore bool

	// Example command : go run echo.go --port 9000 --multicore=true
	flag.IntVar(&port, "port", 9000, "--port 9000")
	flag.BoolVar(&multicore, "multicore", false, "--multicore true")
	flag.Parse()

	// 6
	chat := &chatServer{}
	log.Fatal(gnet.Run(chat, fmt.Sprintf("tcp://:%d", port), gnet.WithMulticore(multicore)))
}
