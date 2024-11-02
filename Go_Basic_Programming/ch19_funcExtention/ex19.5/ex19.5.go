package main

import "fmt"

func main() {
	i := 0

	fn := func() {
		i += 10
	}
	fmt.Println(i) // 0

	i++

	fmt.Println(i) // 1

	fn() // 함수가 호출되는 순간의 i이 참조형으로 함수리터럴에 복사됨.

	fmt.Println(i) // 11

}
