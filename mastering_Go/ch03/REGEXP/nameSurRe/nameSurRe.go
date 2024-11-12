package main

import (
	"fmt"
	"os"
	"regexp"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	// ^, $ 줄의 시작과 끝, 대문자로 시작, 그다음부터는 아무 길이(*)의 소문자로 된 문자열
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func main() {
	args := os.Args

	res := matchNameSur(args[1])
	fmt.Println(res)
}
