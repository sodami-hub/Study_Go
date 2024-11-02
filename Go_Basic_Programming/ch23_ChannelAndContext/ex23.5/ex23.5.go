// select문을 활용한 채널
// 채널에서 데이터가 들어오기를 대기하는 상황에서 만약 데이터가 들어오지 않으면 다른 작업을 하거나,
// 아니면 여러 채널을 동시에 대기하고 싶을 때 어떻게 해야 할까? 바로 select 문을 사용해서 대기한다.

package main

import (
	"fmt"
	"sync"
	"time"
)

func square(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Printf("square : %d\n", n*n)
			time.Sleep(500 * time.Millisecond)
		case <-quit:
			wg.Done()
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	quit := make(chan bool)

	wg.Add(1)
	go square(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}
