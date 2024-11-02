package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	ch := make(chan int) // 1. 채널 생성
	//var ch chan int = make(chan int)

	wg.Add(1)
	go square(&wg, ch) //2. 고루틴 생성
	ch <- 9            //3. 채널에 데이터 넣음
	wg.Wait()          //4. 작업 완료 대기
}

func square(wg *sync.WaitGroup, ch chan int) {
	n := <-ch               // 5. 채널에서 데이터 꺼내기
	time.Sleep(time.Second) //6. 1초 대기
	fmt.Printf("Square : %d\n", n*n)
	wg.Done()
}
