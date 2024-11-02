package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var person01 Person                 // Person타입 구조체 변수 person01을 선언
	var person01Ptr *Person = &person01 // person01Ptr 포인터 변수에 person01의 주소를 대입

	var person02Ptr *Person = &Person{} // person02Ptr 포인터 변수에 Person 구조체를 만들어서 주소값 대입
	// 위의 경우에는 {... 초기화 값} 필드를 초기화 할 수 있다.

	var person03Ptr *Person = new(Person) // new()를 사용한 초기화 // 필드값을 초기화 할 수 없다.

	fmt.Println(person01Ptr, person02Ptr, person03Ptr)
}
