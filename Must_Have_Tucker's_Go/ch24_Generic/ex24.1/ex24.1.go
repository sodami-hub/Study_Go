package main

import (
	"fmt"

	"golang.org/x/exp/constraints" // 타입을 제한해둔 패키지
)

func add2[T int | float64, K int | float64](a T, b K) T {
	return a + T(b)
}
// func add[T (int8|int16|int32|int64|int) | (float32|float64)]  이런식으로 써야 되는 걸 간단하게 정의해 줌 
func add[T constraints.Integer | constraints.Float](a, b T) T {
	return a + b
}

func add3[T any, K int](a T, b K) {
	fmt.Printf("T = %v , K = %v\n", a, b)
}

func main() {
	var a int = 1
	var b int = 2
	fmt.Println(add(a, b))

	var f1 float64 = 3.14
	var f2 float64 = 1.43
	fmt.Println(add(f1, f2))

	fmt.Println(add2(a, f1))

	add3("abc", 12)
}
