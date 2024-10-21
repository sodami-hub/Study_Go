package main

import (
	"fmt"
	"sync"
	"time"
)

// 채널에서 무한히 데이터를 대기하는 for range문(아래 코드)
// 에서 대기를 끝내고 작업을 완료하도록 하기(ex23.4.go)
func square(wg *sync.WaitGroup, ch chan int) {
	for n := range ch { // 2. main함수의 반복문이 끝나도... 데이터를 계속 기다림 // for range 구문을 사용하면 채널에서 데이터를 계속 기다릴 수 있다.
		// 채널을 닫아줘야 된다. !!
		fmt.Printf("Square: %d\n", n*n)
		time.Sleep(time.Second)
	}
	wg.Done() // 4. 실행되지 않음
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go square(&wg, ch) // 메인아니 다른 한개의 고루틴에서 실행되는 내용.
	for i := 0; i < 10; i++ {
		ch <- i * 2 // 메인 고루틴에서 실행되는 내용 - 1. ch 채널에 데이터 넣기
	}
	wg.Wait() // 3. 작업 완료를 기다림
}
