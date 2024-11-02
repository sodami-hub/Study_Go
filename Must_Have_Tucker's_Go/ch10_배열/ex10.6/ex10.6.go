package main

import "fmt"

func main() {
	a := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	for _, arr := range a { // 다차원 배열과 range 순회를 통한 처리
		for _, i := range arr {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
