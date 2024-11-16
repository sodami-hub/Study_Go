// 빈인터페이스를 사용하는 방법은 쉽다. 그저 interface{}를 이요해 모든 타입의 변수를 매개변수로 받을 수 있고 모든 값을 반환할 수 있다.
// 하지만 매우 주의해야 한다. 위험한 상황에 대해서 assertions.go 에서 살펴보겠다.

package main

import "fmt"

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1 int
	F2 S1
}

func Print(s interface{}) {
	fmt.Println(s)
}

func main() {
	v1 := S1{10, "hello"}
	v2 := S2{100, v1}

	Print(v1) // 구조체를 매개변수로 받는다.
	Print(v2)

	Print(123)              // 정수를 매개변수로
	Print("Go is the best") // 문자열을 매개변수로
}
