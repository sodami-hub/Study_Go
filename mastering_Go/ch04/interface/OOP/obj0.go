/*
프라이빗 필드와 함수를 정의하는 방법, 합성, 다형성, 익명 구조체를 기존 구조체에 임베드해서 모든 필드에 접근하는 방법
*/

package main

import (
	"fmt"
)

type IntA interface {
	foo()
}

type IntB interface {
	bar()
}

// IntA, IntB를 만족하는 데이터 타입을 구현하면 IntC도 만족한다.
type IntC interface {
	IntA
	IntB
}

func processA(s IntA) {
	fmt.Printf("%T\n", s)
}

type a struct {
	XX int
	YY int
}

// IntA를 만족한다.
func (varC c) foo() {
	fmt.Println("Foo Processing", varC)
}

// IntB를 만족한다.
func (varC c) bar() {
	fmt.Println("Bar Processing", varC)
}

type b struct {
	AA string
	XX int
}

// 구조체 c는 a,b 타입의 필드를 가지고 있고, IntA, IntB 인터페이스를 구현하기때문에 IntC 인터페이스또한 만족한다.
type c struct {
	A a
	B b
}

// compose 구조체는 a 구조체의 필드를 가진다.
type compose struct {
	field1 int
	a      // 익명 구조체.
}

// 다른 구조체는 같은 이름의 메서드를 가질 수 있다.
func (A a) A() {
	fmt.Println("Function A() for A")
}

func (B b) A() {
	fmt.Println("Function A() for B")
}

func main() {

	// compose 구조체의 사용 예?
	// typeA := a{XX: 1, YY: 2}
	// typeCompose := compose{a: typeA}
	// fmt.Println(typeCompose.a.XX, typeCompose.a.YY)

	var iC c = c{a{120, 12}, b{"-12", -12}}

	iC.A.A()
	iC.B.A()

	// 다음 코드는 동작하지 않는다.
	//iComp := compose{field1: 123, a{567,456}}
	//iComp := compose{field1:123, XX:456, YY:789}
	iComp := compose{123, a{456, 789}}
	// 어떤 구조체에서 익명 구조체를 사용할 때는 익명 구조체의  필드는 iComp.XX, iComp.YY처럼 직접 접근할 수 있다.
	fmt.Println(iComp.XX, iComp.YY, iComp.field1)

	iC.bar()
	processA(iC)
}
