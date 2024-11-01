package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const input = "Now is the winter of our discontent,\n Made glorious summer by this sun of York.\n"

	scanner := bufio.NewScanner(strings.NewReader(input)) // 스캐너생성
	scanner.Split(bufio.ScanWords)                        // 단어 단위로 검색

	count := 0

	for scanner.Scan() {
		fmt.Printf("%s ", scanner.Text())
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input :", err)
	}
	fmt.Printf("%d\n", count)

}
