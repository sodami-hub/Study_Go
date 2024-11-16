/*
타입 단언(type assertion)과 타입 스위치(type switch)

- 빈 인터페이스에 저장된 데이터 타입을 알아내는 방법
*/

package main

import (
	"fmt"
)

type Secret struct {
	SecretValue string
}

type Entry struct {
	F1 int
	F2 string
	F3 Secret
}

// Secret과 Entry 타입을 지원하는 타입스위치
func Teststruct(x interface{}) {
	// 타입 스위치
	switch T := x.(type) {
	case Secret:
		fmt.Println("Secret type")
	case Entry:
		fmt.Println("Entry type")
	default:
		fmt.Printf("Not support type : %T\n", T)
	}
}

// 입력 매개변수의 타입을 출력
func Learn(x interface{}) {
	switch T := x.(type) {
	default:
		fmt.Printf("Data type: %T\n", T)
	}
}

func catchValue(x interface{}) {
	if v, ok := x.(int); ok { // x가 int 형이면 int값과 true를 반환한다. 아니라면 false를 반환한다.
		fmt.Printf("datatype : %T, value : %v\n", x, v)
	} else {
		fmt.Println("int 타입 아님")
	}

}

func main() {
	A := Entry{100, "F2", Secret{"myPassword"}}
	Teststruct(A)
	Teststruct(A.F3)
	Teststruct("A string")

	Learn(12.23)
	Learn('e')

	catchValue(12.3)
	catchValue(1)
}
