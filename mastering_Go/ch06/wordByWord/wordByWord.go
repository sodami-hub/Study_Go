package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func wordByWord(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("file read error", err)
			return err
		}

		re := regexp.MustCompile("[^\\s]+")
		words := re.FindAllString(line, -1)
		for i := 0; i < len(words); i++ {
			fmt.Print(words[i], " ")
		}
		fmt.Println()
	}
	return nil
}

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("1개 이상의 읽을 파일 넣어라")
		return
	}

	for _, v := range args[1:] {
		err := wordByWord(v)
		if err != nil {
			fmt.Println("error!!", err)
			break
		}
	}
}
