// 함수를 반환하거나, 함수를 변수의 값으로 사용

package main

import "fmt"

type funnn func(a, b int) int // 함수를 별칭으로 사용.

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) funnn { // 함수를 반환하는 함수 - 별칭사용
	if op == "+" {
		return add
	} else if op == "*" {
		return mul
	}
	return nil
}

func main() {
	var operator funnn //func(a, b int) int // 함수를 값으로 갖는(주소를 갖는) 변수 - 별칭 사용하지 않음 그러나 별칭 안쓴다고 지랄해서 바꿈
	operator = getOperator("*")

	var result = operator(1, 4)
	fmt.Println(result)
}
