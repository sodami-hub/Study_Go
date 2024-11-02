package main

/*
최소 출력 너비 지정
*/

import "fmt"

func main() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)
	fmt.Printf("%05d, %05d\n", a, b) // 최소 너비보다 짧은 공간을 0으로
	fmt.Printf("%-5d, %-5d\n", a, b) // 최소 너비보다 짧은 수를 왼쪽 정렬

	fmt.Printf("%5d, %5d\n", c, c) // 최소 너비보다 긴 값은 모두 지정한 최소 너비가 무시되어 출력
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-05d, %-05d\n", c, c)
}
