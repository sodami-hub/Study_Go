package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// localhost:8080 에서 리스닝하는 웹 서비스를 구성한다.
var addr = flag.String("listen", "localhost:8080", "listen address")

func main() {
	flag.Parse()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Caddy는 클라이언트의 요청을 이 소켓 주소로 넘겨준다.
	err := run(*addr, c)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server stopped")
}

func run(addr string, c chan os.Signal) error {
	mux := http.NewServeMux()
	mux.Handle("/",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			/*
				Caddy는 어느 클라이언트가 요청을 보냈든 상관없이 웹 서비스에게 요청을 전달한다. 마찬가지로 웹 서비스의 모든 응답은 Caddy에게
				전달되어 다시 올바른 클라이언트에게로 전달된다. Caddy는 어느 클라이언트로부터 요청이 왔는지, 각 요청마다 클라이언트의 IP 주소를
				X-Forwarded-For 헤더 필드에 값을 넣는다.
			*/
			clientAddr := r.Header.Get("X-Forwarded-For")
			log.Printf("%s -> %s -> %s", clientAddr, r.RemoteAddr, r.URL)

			// 핸들러는 HTML 문서를 바이트 슬라이스 포맷으로 응답에 쓴다.
			_, _ = w.Write(index)
		}),
	)

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
	}

	go func() {
		for {
			if <-c == os.Interrupt {
				_ = srv.Close()
				return
			}
		}
	}()

	fmt.Printf("Listening on %s ... \n", srv.Addr)
	err := srv.ListenAndServe()
	if err == http.ErrServerClosed {
		err = nil
	}

	return err
}

/*
리소스에는 http://localhost:2020/style.css 처럼 전체 URL을 사용하지 않는다. 왜냐하면 백엔드 웹 서비스에서는 Caddy에 대한 정보가 없기 때문이며,
또한 클라이언트가 어떻게 Caddy에 접근해야 할지에 대한 정보가 없기 때문이다. 리소스 주소에 스키마 정보와 호스트 네임, 포트 번호를 생략하면 클라이언트의 웹 부라우저가
HTML내의 /style.css라는 리소스를 접했을 때 최초 요청으로부터 스키마 정보와 호스트 네임, 포트번호를 가져와서 사용한다.
이를 위해 Caddy환경 구성에서 요청의 일부는 정적 파일을 서빙하도록 하고 그 외에는 백엔드 웹 서비스로 전달하도록 Caddy를 환경구성해야 한다.
*/
var index = []byte(`<!DOCKTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Caddy Backend Test</title>
	<link href="/style.css" rel="stylesheet">
</head>
<body>
	<h1>Caddy를 사용한 테스트 서버입니다.</h1>
	<p><img src="/hiking.svg" alt="hiking gopher"></p>
</body>
</html>`)
