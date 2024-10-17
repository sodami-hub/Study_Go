package main

import "fmt"

// Go 언어에서는 -er을 붙여서 인터페이스명을 만드는 것을 권장한다.
type Stringer interface {
	String() string
}

type Student struct {
	Name string
	age  int
}

func (s Student) String() string { // Stduent 구조체의 메소드는 Stringer 인터페이스를 구현한 메소드이다.
	return fmt.Sprintf("안녕! 나는 %d살 %s라고 해", s.age, s.Name) // Sprintf 화면에 출력하는 것이 아닌 string 타입으로 반환하는 함수
}

func main() {
	student := Student{"철수", 12}

	stringer := Stringer(student) // 변수 선언과 할당을 한 줄로 합침 - 쉽게 말해서 자바의 업케스팅인듯..
	// Stringer stringer = new Student(); -> Student 클래스가 Stringer인터페이스를 구현했기 때문에 Student를 Stringer객체 생성가능

	str := stringer.String() // Stringer에서 String올 호출한 값은
	fmt.Println(str)
	str = student.String() // Student에서 String을 호출한 값이 같다.
	fmt.Println(str)

	fmt.Printf("%s\n", stringer.String())
}
