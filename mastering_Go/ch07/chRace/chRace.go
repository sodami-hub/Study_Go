package main

import (
	"fmt"
)

func printer(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func main() {
	// 버퍼를 사용하지 않는 채널이다.
	var ch chan bool = make(chan bool)

	// 5개의 값을 하나의 고루틴을 이용해 채널에 쓴다.
	go printer(ch, 5)

	// 중요 : 채널 ch가 닫혔기 때문에
	// range 루프는 스스로 끝난다.
	for val := range ch {
		fmt.Print(val, " ")
	}
	fmt.Println()

	for i := 0; i < 15; i++ {
		fmt.Println(<-ch, " ")
	}
	fmt.Println()
}
