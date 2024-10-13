package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	for { // 1. 무한루프
		fmt.Println("입력하세요.")
		var number int
		_, err := fmt.Scanln(&number) //2. 한줄 입력 받기
		if err != nil {               //2. 입력값이 숫자가 아니면
			fmt.Println("숫자를 입력하세요.")

			stdin.ReadString('\n') //3. 키보드 버퍼를 지운다.
			continue               //5. 1(for문의 시작)로 돌아간다.
		}
		fmt.Printf("입력한 숫자는 %d입니다.\n", number)
		if number%2 == 0 { //6. 짝수 검사를한다.
			break //7. 짝수이면 for문 종료한다.
		}
	}
	fmt.Println("for문이 종료됐습니다.")
}
