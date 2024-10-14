package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	for i, v := range a {
		fmt.Printf("a[%d] = %d  ", i, v)
	}

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d  ", i, v)
	}

	b = a // a배열을 b 변수에 복사 - 배열의 크기가 같아야 된다.

	fmt.Println()

	for i, v := range b {
		fmt.Printf("b[%d] = %d  ", i, v)
	}
}
