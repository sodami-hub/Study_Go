//	멀티 반환 함수 2.
//	변수명을 지정해서 반환하기

package main

import "fmt"

func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return // 리턴값을 명시하지 않아도 선언부의 변수명의 값을 리턴함.
	}
	result = a / b
	success = true
	return
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)
	d, success := Divide(9, 0)
	fmt.Println(d, success)
}
