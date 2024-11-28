package main

import (
	"sync"
)

var wg sync.WaitGroup

// 채널에 아무것도 안하고, 열고 닫기만 하지만 패닉 발생 x
func main() {
	// make를 사용하면 해당 채널을 해당 타입의 0 값으로 초기화한다.
	c := make(chan string)
	// var c chan string  - 이건 패닉 발생 // 이러한 형태는 초기화 하지 않고 선언만 한 형태이다.
	wg.Add(1)
	go func() {
		close(c)
		defer wg.Done()
	}()
	wg.Wait()
}
