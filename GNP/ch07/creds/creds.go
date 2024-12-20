package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"os/user"
	"path/filepath"

	"github.com/sodami-hub/Study_Go/GNP/ch07/auth"
)

// 커맨드 라인에서 그룹 이름을 받기위한 usage 설정 -> init()패키지 초기화 함수 main()보다 먼저 실행된다.
func init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n\t%s <group names>\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

// 그룹 이름을 그룹 ID로 파싱하기
func parseGroupNames(args []string) map[string]struct{} {
	groups := make(map[string]struct{})

	for _, arg := range args {
		grp, err := user.LookupGroup(arg)
		if err != nil {
			log.Println(err)
			continue
		}

		groups[grp.Gid] = struct{}{} // 접속이 허용된 그룹id를 키값으로 빈 구조체와 함께 넣는다.
	}
	return groups
}

// 피어의 인증 정보를 기반으로 사용자 인증
func main() {
	flag.Parse()

	groups := parseGroupNames(flag.Args())
	socket := filepath.Join(os.TempDir(), "creds.sock")
	serverAddr, err := net.ResolveUnixAddr("unix", socket)
	if err != nil {
		log.Fatal(err)
	}

	s, err := net.ListenUnix("unix", serverAddr)
	if err != nil {
		log.Fatal(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		_ = s.Close
	}()

	fmt.Printf("Listening on %s ... \n", socket)

	for {
		conn, err := s.AcceptUnix()
		if err != nil {
			break
		}
		if auth.Allowed(conn, groups) {
			_, err = conn.Write([]byte("Welcome\n"))
			if err == nil {
				// 여기에 연결을 처리하는 고루틴
				continue
			}
		} else {
			_, err = conn.Write([]byte("Access denied\n"))
		}
		if err != nil {
			log.Println(err)
		}
		_ = conn.Close()
	}
}
