package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//#v  포맷 지정자 -> 변수를 Go 문법에 맞는 형태로 출력한다.
	// 변수의 타입과 값을 자세히 보여준다.
	fmt.Printf("%#v\n", wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			fmt.Println(x)
		}(i)
	}

	fmt.Printf("%#v\n", wg)
	wg.Wait()
	fmt.Printf("%#v\n", wg)
	fmt.Println("exiting...")
}
