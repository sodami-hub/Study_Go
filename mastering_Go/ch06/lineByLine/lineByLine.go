package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("읽을 파일을 실행 인수로 넣어라")
		return
	}

	filePath := arguments[1]

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("file open error", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("file read error", err)
			break
		}
		fmt.Println(line)
	}
}
