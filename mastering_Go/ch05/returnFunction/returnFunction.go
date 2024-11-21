package main

import "fmt"

// int형 매개변수를 갖고 int 형 값을 반환하는 함수를 반환하는 funRet 함수
func funRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	}
	return func(k int) int {
		return k * k
	}
}

func main() {
	n := 10
	// n 과 -4를 funRet()에서 반환할 익명 함수를 결정하는 데 활용한다.
	i := funRet(n)
	j := funRet(-4)

	fmt.Printf("%T\n", i)
	fmt.Printf("%T %v\n", j, j)
	fmt.Println("j", j, j(-5))
	fmt.Println("i", i, i(5))

	// 입력 매개변수는 같지만 i와 j는 다른 익명 함수를 가리키고 있다.
}
