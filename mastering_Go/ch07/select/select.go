package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	time.Sleep(time.Second)
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
			//return
		// end에서 값을 받는 case 에서 return을 빼먹었더라도 프로그램이 정상종료될 수 있도록 하는 기능을 추가했다.
		// 4초뒤에 시그널을 보낸다!!
		// time.After()는 해당 블럭이 실행되고부터의 시간을 측정한다.
		case <-time.After(4 * time.Second):
			fmt.Println("time.After()!")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup

	rand.Seed(time.Now().Unix())

	createNumber := make(chan int)
	end := make(chan bool)
	n := 10

	wg.Add(1)
	go func() {
		gen(0, 2*n, createNumber, end)
		wg.Done()
	}()

	for i := 0; i < n; i++ {
		fmt.Print(" ", <-createNumber)
	}

	end <- true
	wg.Wait()

}
