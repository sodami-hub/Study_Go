/*
tftp 서버를 시작하려면 페이로드로 사용할 파일과 요청 수신을 대기할 주소, 두 가지를 전달해 주어야 한다.
*/

package main

import (
	"flag"
	"io/ioutil"
	"log"

	"GNP/ch06/main/tftp"
)

var (
	address = flag.String("a", "127.0.0.1:69", "listen address")
	payload = flag.String("p", "/home/sodami/Documents/Work_Space/study_go/GNP/ch06/main/tftp/image.jpg", "file to serve to client")
)

func main() {
	flag.Parse()

	p, err := ioutil.ReadFile(*payload)
	if err != nil {
		log.Fatal(err)
	}

	s := tftp.Server{Payload: p}
	log.Fatal(s.ListenAndServe(*address))
}
