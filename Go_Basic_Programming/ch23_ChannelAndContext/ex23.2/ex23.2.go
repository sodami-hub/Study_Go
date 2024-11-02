package main

import (
	"fmt"
	//"sync"
)

func main() {
	ch := make(chan int) //1. 크기 0인 채널 생성

	ch <- 9 //2. main() 함수가 여기서 멈춤 - 데이터를 보관할 곳이 없기 때문에 다른 고루틴에서 데이터를 빼갈때까지 영원히 대기
	//	-> deadlock 상태에 빠지게 되고 강제 종료

	/* 주석안에 내용은 정상 처리되는 내용.
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		val := <-ch
		wg.Done()
		fmt.Printf("%v\n", val)
	}()

	ch <- 9

	wg.Wait()
	*/

	fmt.Println("never print") //3. 실행되지 않음
}
