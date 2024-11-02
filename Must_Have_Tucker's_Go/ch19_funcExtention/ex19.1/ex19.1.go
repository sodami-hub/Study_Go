// 가변 인수를 받는 함수
// ... 키워드

package main

import (
	"fmt"
)

func sum(nums ...int) int {
	sum := 0

	fmt.Printf("nums 타입 : %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

// 여러 타입의 인수를 섞어 사용하는 함수
// 인터페이스 사용
func printer(args ...interface{}) {
	for _, v := range args {
		fmt.Print(v, " ")
		switch v.(type) { // 인자의 타입별로 다른 처리가능		// 인터페이스.(type) 인터페이스의 타입 반환
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("another type")
		}
	}
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(100, 200))
	fmt.Println(sum())

	printer("a", 123, "hello", 3.14)

	fmt.Println()

	a := 3.14

	fmt.Printf("%T\n", a)
}
