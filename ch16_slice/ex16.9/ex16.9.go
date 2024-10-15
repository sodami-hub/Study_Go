package main

import "fmt"

// 슬라이스 복사하는 방법 3 ( 서로 영향 주지 않는 새로운 슬라이스로.. ) copy(dest,src) int 사용.

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10) // len 3, cap 10
	slice3 := make([]int, 10)    // len 10, cap 10

	cnt1 := copy(slice2, slice1)
	cnt2 := copy(slice3, slice1)

	fmt.Println(cnt1, slice2)
	fmt.Println(cnt2, slice3)
}
