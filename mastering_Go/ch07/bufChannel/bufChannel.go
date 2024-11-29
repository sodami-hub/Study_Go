package main

import (
	"fmt"
)

func main() {
	// numbers 채널의 크기가 5인 버퍼채널이므로 5개보다 많은 정수를 저장할 수 없다.
	numbers := make(chan int, 5)

	counter := 10

	for i := 0; i < counter; i++ {
		select {
		// 여기서 처리가 일어난다.
		// numbers에 데이터를 넣기시작한다. 하지만 채널이 전부 채워지면 더 이상 데이터를 넣을 수 없기 때문에 default 브랜치가 실행된다.
		case numbers <- i * i:
			fmt.Println("about to process", i)
		default:
			fmt.Print("no space for ", i, "")
		}
	}
	fmt.Println()

	for {
		// 위와 비슷한 방식으로 데이터를 읽는다. 채널의 모든 데이터를 읽고 default 브랜치가 실행되고 return에 의해 종료된다.
		select {
		case num := <-numbers:
			fmt.Print("*", num, " ")
		default:
			fmt.Println("nothing left to read!")
			return
		}
	}
}
