package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"time"

	"ch09/handlers"
	"ch09/middleware"
)

// HTTP/2 서버는 TLS 지원을 필요로 하기에, 이를 위해 매개변수로 인증서의 경로(cert)와 개인키의 경로(pkey)를 전달해야 된다.
// 둘 중 하나의 값이 전달되지 않으면 서버는 일반 HTTP 연결을 대기한다.
// 그리고 run 함수에 커맨드 라인의 플래그 값을 전달한다.
var (
	addr  = flag.String("listen", "localhost:8080", "listen address")
	cert  = flag.String("cert", "", "certificate")
	pkey  = flag.String("key", "", "private key")
	files = flag.String("files", "./files", "static file directory")
)

func main() {
	flag.Parse()

	err := run(*addr, *files, *cert, *pkey)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("server gracefully shutdown")
}

func run(addr, files, cert, pkey string) error {
	mux := http.NewServeMux()
	mux.Handle("/static",
		http.StripPrefix("/static/",
			middleware.RestrictPrefix(".", http.FileServer(http.Dir(files))),
		),
	)
	mux.Handle("/",
		handlers.Methods{
			http.MethodGet: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					if pusher, ok := w.(http.Pusher); ok {
						targets := []string{
							"/static/style.css",
							"/static/hiking.svg",
						}
						for _, target := range targets {
							if err := pusher.Push(target, nil); err != nil {
								log.Printf("%s push failed: %v", target, err)
							}
						}
					}
					http.ServeFile(w, r, filepath.Join(files, "index.html"))
				},
			),
		},
	)

	mux.Handle("/2",
		handlers.Methods{
			http.MethodGet: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					http.ServeFile(w, r, filepath.Join(files, "index2.html"))
				},
			),
		},
	)

	srv := &http.Server{
		Addr:              addr,
		Handler:           mux,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 30 * time.Second,
	}

	done := make(chan struct{})

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)

		for {
			if <-c == os.Interrupt {
				if err := srv.Shutdown(context.Background()); err != nil {
					log.Printf("shutdown: %v", err)
				}
				close(done)
				return
			}
		}
	}()

	log.Printf("Serving files in %q over %s\n", files, srv.Addr)

	var err error
	if cert != "" && pkey != "" {
		log.Println("TLS enabled")
		err = srv.ListenAndServeTLS(cert, pkey)
	} else {
		err = srv.ListenAndServe()
	}

	if err == http.ErrServerClosed {
		err = nil
	}

	<-done

	fmt.Println("error : ", err)
	return err
}
