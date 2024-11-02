package main

import "fmt"

func main() {
	var char rune = '한'

	fmt.Printf("%T\n", char) // char 변수의 타입을 %T 를 이용해서 출력 -> int32
	fmt.Println(char)        // char 값을 출력! int32이기 때문에 정수값으로 출력됨
	fmt.Printf("%c\n", char) // %c 를 이용해서 문자로 출력한다.
}
