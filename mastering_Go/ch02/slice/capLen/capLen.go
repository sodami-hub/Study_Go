package main

import (
	"fmt"
)

func main() {
	// 길이만 정의 됐으므로 길이 = 용량 이다.
	a := make([]int, 4)
	fmt.Println(a, len(a), cap(a))

	b := []int{1, 2, 3, 4, 5}
	fmt.Println(b, len(b), cap(b))

	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	// len == cap이 된 슬라이스에 데이터를 추가하자 cap 이 두배가 됨
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice, len(aSlice), cap(aSlice))

	c := make([]int, 0, 5)
	fmt.Println(c, len(c), cap(c))

}
