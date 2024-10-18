// 파일 핸들을 내부 상태로 사용하는 예

package main

import (
	"fmt"
	"os"
)

type Write func(string)

func writeHello(write Write) { // writeHello에 외부에서 로직을 주입 -> 의존성 주입
	write("Hello World")
}

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("file create fail")
		return
	}
	defer fmt.Println("파일을 닫았습니다.")
	defer f.Close()

	writeHello(func(msg string) {
		fmt.Fprintln(f, msg) // 함수 리터럴에서 외부변수 f사용.
	})
}
