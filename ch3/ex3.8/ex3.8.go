package main

// Scanln을 반복적으로 사용할 때는 아래 주석처리 된 부분처럼 스트림을 비워줘야된다.
// 그렇게 하지않으면 스트림에 문자열이 남아서 계속 에러가 난다.(현재 코드)
/*
hello 3  - > hello를 입력하고..
expected integer -> 에러발생
expected integer -> 다음 입력에도 hello의 문자열이 남아서(e) 에러가 자동 발생
*/
import (
	//"bufio"
	"fmt"
	//"os"
)

func main() {
	//hstdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		//		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		//		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}
