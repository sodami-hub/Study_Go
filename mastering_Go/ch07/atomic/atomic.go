package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 원자적인 int64 값을 갖고 있는 구조체
type atomCounter struct {
	val int64
}

// atomic.LoadInt64()를 이용해 현재의 int64 변수 값을 원자적으로 가져오는 함수다.
func (c *atomCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	X := 100
	Y := 4

	var wg sync.WaitGroup
	counter := atomCounter{}
	for i := 0; i < X; i++ {
		wg.Add(1)
		go func(no int) {
			defer wg.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println(counter.Value())
}
