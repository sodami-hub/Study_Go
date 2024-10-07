package main

import "fmt"

func main() {
	var a int
	var b int

	//scanf 정숫값 두개를 입력받은 서식에 맞게 두 숫자 사이에 공란이 와야 됨
	n, err := fmt.Scanf("%d %d\n", &a, &b)
	if err != nil {
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
