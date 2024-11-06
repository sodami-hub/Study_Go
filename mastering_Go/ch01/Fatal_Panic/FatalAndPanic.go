package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 1 {
		log.Fatal("Fatal") // 메세지 출력(log.Print()) 후 os.Exit(1) 을 호출해서 프로그램 전체 종료
	}
	log.Panic("Panic") // 메세지 출력 이후, 함수가 호출된 곳으로 되돌아 감 여기서는 main.main()
}
