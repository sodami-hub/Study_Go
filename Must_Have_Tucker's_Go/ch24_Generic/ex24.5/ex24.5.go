package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Map[F, T any](s []F, f func(F) T) []T {
	rst := make([]T, len(s))
	for i, v := range s {
		rst[i] = f(v)
	}
	return rst
}

func main() {
	//각 값을 두배씩 증가시키는 슬라이스
	doubled := Map([]int{1, 2, 3}, func(v int) int {
		return v * 2
	})

	//대문자로 변경하는 슬라이스
	uppered := Map([]string{"hello", "world", "abcde"}, func(v string) string {
		return strings.ToUpper(v)
	})

	// 문자열로 변경하는 슬라이스
	toString := Map([]int{1, 2, 3}, func(v int) string {
		return "str" + strconv.Itoa(v)
	})

	fmt.Println(doubled)
	fmt.Println(uppered)
	fmt.Println(toString)
}
