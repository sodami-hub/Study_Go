package main

import (
	"errors"
	"fmt"
)

func Atoi(input string) (int, error) {
	var res int
	res = 0
	for _, v := range input {
		if v >= '0' && v <= '9' {
			res *= 10
			res += int(v - '0')
		} else {
			return 0, errors.New("숫자형식의 문자열을 입력해야 됩니다. input :" + input)
		}
	}
	return res, nil
}

func main() {
	result, err := Atoi("123a")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
}
