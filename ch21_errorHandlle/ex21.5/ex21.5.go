//type interface{} any
//func panic(v any)
//panic() 내장 함수의 인수로 interface{} 타입 즉 모든 타입을 사용할 수 있다.

package main

import "fmt"

func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없다.")
	}
	fmt.Println(a / b)
}

func main() {
	divide(9, 3)
	divide(9, 0)
}
