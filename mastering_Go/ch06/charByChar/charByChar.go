package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func charByChar(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}

		// 한줄에 한문자만 출력
		// char는 rune 타입이므로 string으로 형변환해야 된다.
		for _, char := range line {
			fmt.Print(string(char))
			break
		}
		fmt.Println()
	}
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("읽어들일 파일을 하나 이상 넣으시오")
		return
	}

	for _, file := range args[1:] {
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
