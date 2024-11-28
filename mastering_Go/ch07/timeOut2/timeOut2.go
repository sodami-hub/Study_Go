package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var result = make(chan bool)

func timeout(t time.Duration) {
	temp := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		// close() 채널을 닫는 함수
		defer close(temp)
	}()

	// 타임아웃시간이 고루틴에서 채널을 닫는 시간보다 길면 result에 false가 보내진다.
	// 타임아웃이 고루틴에서 채널을 닫는 시간보다 짧으면 true가 발생해서 타임아웃이 발생한다.
	select {
	case <-temp:
		result <- false
	case <-time.After(t):
		result <- true
	}
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("추가 실행인수 필요")
		return
	}

	t, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	duration := time.Duration(int32(t)) * time.Millisecond
	fmt.Printf("timeout period is %s\n", duration)

	go timeout(duration)

	val := <-result

	if val {
		fmt.Println("Time out")
	} else {
		fmt.Println("ok")
	}
}
