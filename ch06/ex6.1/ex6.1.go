package main

import "fmt"

func main() {
	const C int = 10
	fmt.Println(C)
	/*
		var b int = C * 20
		C = 20 // error 상수는 대입문의 좌변에 올 수 없다
		fmt.Println(&C) // error 상수의 메모리 주솟값에는 접근할 수 없다.
	*/
}
