package main

import "fmt"

type Stringer interface {
	String() string
}

type Student struct {
	Age int
}

// 다형성을 말하는 부분.
// 구조체(클래스) a가 인터페이스 i의 구현 메소드를 가지고 있다면 구조체(클래스) a는 i타입으로 그리고 다시 a타입으로 변환할 수 있다.

func (s *Student) String() string { // 1. Student는 Stringer인터페이스의 구현체
	return fmt.Sprintf("Student Age : %d", s.Age)
}

func PrintAge(stringer Stringer) { // 3. Student 객체 s를 Stringer형으로 자동 타입변환시킴.
	s := stringer.(*Student) //4. Stringer 로 자동 타입변환된 s를 다시 Student 객체로 타입변환... s는 -> s instanceof Student
	fmt.Printf("Age: %d\n", s.Age)
}

func main() {
	s := &Student{15} // 2. s는 Student 객체

	PrintAge(s)
}
