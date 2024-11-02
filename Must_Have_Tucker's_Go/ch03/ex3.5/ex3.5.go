package main

import "fmt"

func main() {
	var a int
	var b int

	//scan 표준 입력에서 값을 입력 받음
	n, err := fmt.Scan(&a, &b) // n은 성공적으로 입력한 값의 갯수, err은 입력시 발생한 에러를 반환한다.
	if err != nil {            // err 이 nil(값이 없다)이 아니면 error이다.
		fmt.Println(n, err)
	} else {
		fmt.Println(n, a, b)
	}
}
