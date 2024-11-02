package main

import "fmt"

func main() {
	var a int = 500
	var p *int // 포인터 변수의 기본값은 nil 이다. nil이 아니면 유효한 메모리 주소를 가지고 있다는 의미이다.

	p = &a
	fmt.Printf("p의 값 : %p\n", p)            // 메모리 주솟값 출력
	fmt.Printf("p가 가리키는 메모리의 값 : %d\n", *p) // p가 가리키는 메모리의 값을 출력

	*p = 100 // p가 가리키는 메모리에 저장된 값을 변경
	fmt.Printf("a의 값 : %d", a)

}
