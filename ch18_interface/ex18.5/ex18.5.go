package main

import "fmt"

// 빈 인터페이스는 어떠한 값이라도 받을 수 있는 함수, 메서드 변숫값을 만들 때 사용한다.

func PrintVal(v interface{}) { // 빈 인터페이스를 인수로 받는 함수
	switch t := v.(type) { // v의 타입에 따라서 다른 로직 수행
	case int:
		fmt.Printf("v is int %d\n", int(t))
	case float64:
		fmt.Printf("v is float64 %.3f\n", float64(t))
	case string:
		fmt.Printf("v is string %s\n", string(t))
	default:
		fmt.Printf("Not supported type : %T:%v\n", t, t)
	}
}

type Student struct {
	Age int
}

func main() {
	PrintVal(10)
	PrintVal(3.14)
	PrintVal("hello")
	PrintVal(Student{15})
}
