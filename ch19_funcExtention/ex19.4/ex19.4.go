// 함수 리터럴을 사용한 함수 반환

package main

import "fmt"

type opFunc func(a, b int) int

func getOperator(op string) opFunc {
	if op == "+" {
		return func(a, b int) int { // 함수 리터럴을 사용해서 함수를 정의하고 반환 opFunc의 형식과 동일해야 됨.
			return a + b
		}
	} else if op == "-" {
		return func(a, b int) int {
			return a - b
		}
	}
	return nil
}

func main() {

	a := getOperator("-")

	result := a(4, 3)

	fmt.Println(result)

	// 함수 리터럴의 호출 방법

	fn := func(a, b int) int {
		return a + b
	}
	result = fn(3, 4)
	fmt.Println(result)

	result = func(a, b int) int {
		return a + b
	}(3, 4)
	fmt.Println(result)
}
