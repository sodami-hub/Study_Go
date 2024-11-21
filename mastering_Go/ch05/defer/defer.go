package main

import "fmt"

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	j := 0
	for i := 3; i > 0; i-- {
		j++
		// 익명함수는 클로저이므로 함수 범위 밖의 변수에 접근할 수 있다.
		defer func() {
			fmt.Print(j, " ")
		}()
	}
	fmt.Println()
}

// defer를 가장 바람직하게 사용한 함수, 원하는 변수를 명시적으로 익명 함수에 전달해 이해하기 쉽다.
func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
}
