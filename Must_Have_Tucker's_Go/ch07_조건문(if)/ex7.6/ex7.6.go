// if 초기문; 조건문 {}
package main

import "fmt"

func getmyage() (int, bool) {
	return 33, true
}

func main() {

	if age, ok := getmyage(); ok && age < 20 {
		fmt.Println("very young")
	} else if ok && age < 30 {
		fmt.Println("young")
	} else if ok {
		fmt.Println("you are beautiful", age)
	} else {
		fmt.Println("error")
	}

	// fmt.Println("your age is", age) // age는 소멸됨 -> if 블럭스코프를 갖는 지역변수
}
