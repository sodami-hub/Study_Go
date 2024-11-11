// 두 개의 배열을 연결해 새로운 슬라이스를 만드는 함수

package main

import (
	"fmt"
)

func arraysToSlice(a, b [5]int) []int {
	var slice []int
	slice = append(slice, a[:5]...)
	slice = append(slice, b[:5]...)
	return slice
}

func arraysToArray(a, b [5]int) [10]int {
	var array [10]int

	for i := 0; i < 10; i++ {
		if i < 5 {
			array[i] = a[i]
		} else {
			array[i] = b[i-5]
		}
	}
	return array
}

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{6, 7, 8, 9, 10}

	slice := arraysToSlice(a, b)
	fmt.Printf("type : %T, %v\n", slice, slice)

	array := arraysToArray(a, b)
	fmt.Printf("type : %T, %v\n", array, array)
}
