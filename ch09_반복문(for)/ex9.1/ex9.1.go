package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ { // for 초기값; 조건문 ; 후처리
		fmt.Print(i, " ")
	}

	fmt.Println()

	i := 0

	for true { // for true 무한루프 -> true 생략가능(switch-case 와 동일)
		fmt.Print(i, " ")
		i++
		if i == 10 {
			break
		}
	}
}
