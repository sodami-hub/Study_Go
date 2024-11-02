// switch 초기문; 비굣값 {}
package main

import "fmt"

func getMyAge() int {
	return 33
}

func main() {
	switch age := getMyAge(); age {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 3")
	default:
		fmt.Println("My age is", age)
	}

	// fmt.Println("my age is", age) // age는 소멸됨 -> age는 블럭스코프를 갖는 지역변수
}
