package main

import (
	"fmt"
	"sync"
)

// writeToChannel() 함수는 c 채널에 x를 보내고 채널을 닫는다
func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

// 채널 ch는 bool 타입이고 printer()는 ch 채널에 true를 보낸다.
func printer(ch chan bool) {
	ch <- true
}

func main() {
	// 이 채널은 1의 크기를 갖는 버퍼 채널이다.
	//따라서 해당 버퍼가 채워지는 순간 채널을 닫을 수 있고 고루틴은 다음 구문의 실행을 계속하고 반환할 수 있다.
	// 버퍼를 사용하지 안흔ㄴ 채널은 다른 동작을 한다. 버퍼를 사용하지 않는 채널에 값을 전달하면 다른 누군가가 값을 가져갈 때까지 실행을 멈춘다.
	// 이 코드에서는 실행을 멈추는 것을 원하지 않기 때문에 버퍼 채널을 사용해야 한다.
	c := make(chan int, 1)

	var wg sync.WaitGroup

	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)

	fmt.Println("Read :", <-c)

	_, ok := <-c
	if ok {
		fmt.Println("channel is open")
	} else {
		fmt.Println("channel is closed")
	}

	wg.Wait()

	var ch chan bool = make(chan bool)

	for i := 0; i < 5; i++ {
		// 버퍼를 사용하지 않는 채널을 만든 뒤 아무런 동기화 없이 5개의 고루틴을 만들었다.
		go printer(ch)
	}

	// 채널과 range
	// 중요: ch 채널이 닫히지 않았기 때문에
	// range 루프는 스스로 끝나지 않는다.
	n := 0
	for i := range ch {
		fmt.Println(i)
		if i == true {
			n++
		}

		if n > 2 {
			fmt.Println("n :", n)
			close(ch)
			break
		}
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

}
