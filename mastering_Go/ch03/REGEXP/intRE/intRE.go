package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchInt(s string) bool {
	t := []byte(s)
	// +,- 로 시작(생략가능 '?'), 1개 이상의 아무길이(+)의 숫자(\d)로 이루어진 문자열
	re := regexp.MustCompile(`^[+-]?\d+$`)
	return re.Match(t)
}

func main() {
	args := os.Args

	res := matchInt(args[1])
	fmt.Println(res)
}
