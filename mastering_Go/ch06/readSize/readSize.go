package main

import (
	"fmt"
	"io"
	"os"
)

func readSize(f *os.File, size int) []byte {
	// 매개변수의 size만큼의 버퍼를 만든다.
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return buffer[:n]
}
func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("1개 이상의 읽을 파일 넣어라")
		return
	}

	for _, v := range args[1:] {
		f, err := os.Open(v)
		read := readSize(f, 1024)
		if err != nil {
			fmt.Println("error!!", err)
			break
		}
		fmt.Println(string(read))
	}
}
