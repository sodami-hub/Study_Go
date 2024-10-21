package main

import (
	"fmt"
	"time"
)

func PrintHangul() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}
	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c ", v)
	}
}

func PrintNumber() {
	for i := 1; i < 10; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go PrintHangul() // go 키워드를 쓰고 함수를 호출하면 새로운 고루틴을 생성
	go PrintNumber()

	time.Sleep(3 * time.Second) // 기다리지 않으면 다른 두개의 고루틴도 즉시 종료됨. 아무것도 출력되지 않음
}
