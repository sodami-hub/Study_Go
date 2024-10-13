package main

import "fmt"

func main() {
	for i := range 10 { // 정숫값의 순회 0~9까지 순회함
		fmt.Print(i, " ")
	}
}
