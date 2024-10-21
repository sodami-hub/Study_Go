// 일정 간격으로 실행
// time 패키지의 Tick()함수로 원하는 시간 간격으로 신호를 보내주는 채널을 만들 수 있다.

package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)            // 1초 간격 시그널 - time.Tick() 일정 간격으로 현재 시각을 나타내는 Time 객체를 반환
	terminate := time.After(10 * time.Second) // 10초 이후 시그널 - 일정 시간 경과 후에 현재 시각을 나타내는 Time 객체를 반환

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-terminate:
			fmt.Println("Terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Printf("square :%d\n", n*n)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan int)
	wg.Add(1)
	go square(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
