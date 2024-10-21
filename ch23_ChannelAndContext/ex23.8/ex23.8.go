package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	// 1. 취소 가능한 컨텍스트 생성
	// 상위 컨텍스트를 인수로 넣으면 그 컨텍스트를 감싼 새로운 컨텍스트를 만들어준다.
	// 상위 컨텍스트가 없다면 가장 기본적인 컨텍스트인 context.Background()를 넣어준다.
	// context.WithCancel 함수는 값을 두개 반환한다. 첫번째가 컨텍스트 객체, 두번째가 취소 함수이다.
	// 취소 함수를 사용해서 원할 때 취소할 수 있다.
	ctx, cancel := context.WithCancel(context.Background())

	//	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) // 3초뒤에 작업을 종료하는 컨텍스트

	go PrintEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel() // 5초 뒤 취소 - 취소함수 사용. 컨텍스트의 Done() 채널에 시그널을 보내 작업이 취소될 수 있도록 한다.

	wg.Wait()
}

func PrintEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done(): // 컨텍스트 Done() 채널의 시그널을 검사.
			wg.Done()
			return
		case <-tick:
			fmt.Println("Tick")
		}
	}
}
