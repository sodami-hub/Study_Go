package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func add(c chan int) {
	sum := 0
	t := time.NewTimer(2 * time.Second)

	for {
		select {
		case input := <-c:
			sum = sum + input
		case <-t.C:
			// 채널 c가 nil이 되면 send()에서 더이상 데이터를 보내지 않는다.
			c = nil
			fmt.Println(sum)
			wg.Done()
		}
	}
}

// 지속적으로 c 채널로 난수를 본낸다.
func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	rand.Seed(time.Now().Unix())
	wg.Add(1)
	go add(c)
	go send(c)
	wg.Wait()
}
