package main

// Scanln을 반복적으로 사용할 때는 아래 주석처리 된 부분처럼 스트림을 비워줘야된다.
// 그렇게 하지않으면 스트림에 문자열이 남아서 계속 에러가 난다.(현재 코드)
/*
hello 3  - > hello를 입력하고..
expected integer -> 에러발생
expected integer -> 다음 입력에도 hello의 문자열이 남아서(e) 에러가 자동 발생
*/
import (
	//"bufio"	// io를 담당하는 패키지
	"fmt"
	//"os"		// 표준 입출력 등을 가지고 있는 패키지
)

// bufio는 입력 스트림으로부터 한 줄을 읽는 Reader 객체를 제공한다.

func main() {
	// 표준 입력을 읽는 객체, NewReader() 함수는 인수로 입력되는 입력 스트림을 가지고 Reader 객체를 생성한다.
	//hstdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)

	if err != nil {
		fmt.Println(err)
		// 줄 바꿈 문자가 나올 때까지 읽는다. 이렇게 하면 표준 입력 스트림이 비워진다.
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
