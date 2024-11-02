package main

import (
	"fmt"
)

func Atoi(str string) (int, error) {
	rst := 0
	for _, r := range str {
		if r >= '0' && r <= '9' {
			rst *= 10
			rst += int(r - '0')
		} else {
			return 0, fmt.Errorf("숫자만 입력하세요. 문자 : %c", r)
		}
	}
	return rst, nil
}

func main() {
	n, err := Atoi("34cd")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n)
	}
}
