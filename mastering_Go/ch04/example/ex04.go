// 빈 인터페이스를 받아 직접 만든 두 개의 구조체를 구분하는 함수를구현

package main

import "fmt"

type form01 struct {
	intA int
	intB int
}

type form02 struct {
	stringA string
	stringB string
}

func findType(i interface{}) {
	switch i.(type) {
	case form01:
		fmt.Println("this is a form01")
	case form02:
		fmt.Println("this is a form02")
	default:
		fmt.Println("what?!")
	}
}

func main() {
	a := form01{1, 1}
	b := form02{"일", "일"}

	findType(a)
	findType(b)
}
