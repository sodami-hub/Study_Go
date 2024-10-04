package main

import "fmt"

func main() {
	var a int
	var b int

	//scanln 한 줄을 입력받아서 공란으로 구분된 값들을 읽는다.
	n, err := fmt.Scanln("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
