package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("error")
		return
	}

	// 3~5 defer로 함수 종료 전에 반드시 호출되어야 할 코드를 지정한다. 출력결과 역순으로 호출됐다.
	// 밑에까지 다 실행하고 다시 올라오면 서 실행한다.
	defer fmt.Println("반드시 호출 됨") // 3
	defer f.Close()               // 4
	defer fmt.Println("파일을 닫았다.") // 5

	fmt.Println("파일에 hello world를 쓴다.")
	fmt.Fprintln(f, "Hello world") // 파일에
}
