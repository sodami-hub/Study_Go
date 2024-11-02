package main

import (
	"fmt"
	"sync"
	"time"
)

// 데이터를 모두 넣고 채널이 더는 필요없을 때 close(ch) 호출해서 채널을 닫아주면, for range 에서 데이터를 모두 처리하고 채널이 닫히면 for range문도 종료된다.
func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(500 * time.Millisecond)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch) // 채널 닫음!

	wg.Wait()
}
