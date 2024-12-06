package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	// filepath.Base() 함수는 경로의 마지막 값을 반환한다. 예를들어 os.Args[0]이 매개변수로 주어지면 filepath.Base()함수는
	// 경로의 마지막 값인 파일의 이름을 반환한다.
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s URL \n", filepath.Base(os.Args[0]))
		return
	}

	URL := os.Args[1]
	data, err := http.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 서버 응답의 본문이 들어 있는 data.Body에서 데이터를 읽어 os.Stdout에 출력한다.
	// os.Stdout은 항상 열려있다. 여기서 모든 데이터는 표준 출력에 쓴다.(일반적으로 터미널 화면)
	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	data.Body.Close()
}
