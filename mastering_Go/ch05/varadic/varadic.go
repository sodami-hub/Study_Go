package main

import (
	"fmt"
)

func addFloats(message string, s ...float64) float64 {
	fmt.Println(message)
	var sum float64 = 0
	for _, a := range s {
		sum += a
	}
	s[0] = -1000
	return sum
}

func everything(input ...interface{}) {
	fmt.Println(input)
}

func main() {
	sum := addFloats("Adding numbers...", 1.2, 1.5, 2.1, 4.6, -20, 123)

	fmt.Println("sum :", sum)
	s := []float64{1.1, 2.2, 3.3, 4.4}

	// 슬라이스와 언팩 연산자
	sum = addFloats("Adding numbers...", s...)
	fmt.Println("sum: ", sum)
	fmt.Println(s)
	everything(s) // 슬라이스를 ...s 과 같이 언팩하지 않았기 때문에 동작한다.

	// []string을 바로 []interface{} 형태로 전달할 수 없다.
	// 전달하기 전에 먼저 변환해야 한다.
	empty := make([]interface{}, len(s))
	for i, v := range s {
		empty[i] = v
	}
	// 언팩이 가능하다.
	everything(empty...)

}
