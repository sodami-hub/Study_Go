package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

/*
WithCancel() 메서드는 부모 컨텍스트의 복사본과 새로운 Done 채널을 반환한다.
함수로 된 cancel변수는 context.CancelFunc()의 값이다. context.WithCancel()함수는 기존 Context를 사용해
캔슬레이션으로 그 자식을 생성한다.
Done 채널이 닫히는 시점은 cancel()함수를 호출할 때나 부모 컨텍스트의 Done 채널이 닫힐 때다.
*/

func f1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()

	// 4초가 지나면 고루틴에서 ㅋ컨텍스르트를 취소하고 c1.Done() 채널에 시그널을 보낸다.
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	// t가 3인경우 함수는 3초 후에 time.After에서 신호를 받아서 메시지를 출력하고 종료된다.
	// 컨텍스트가 취소되기 전에 함수가 종료되므로 c1.Done() 채널에서 신호를 받지 않는다.
	select {
	case <-c1.Done():
		fmt.Println("f1() Done:", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1():", r)
	}

	return
}

func f2(t int) {
	c2 := context.Background()
	// WithTimeout() time.Duration으로 설정된 시간이 경과하면 cancel() 함수가 호출
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("f2() Done:", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f2():", r)
	}
	return
}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	// deadline이 지나면 자동으로 cancel()이 호출된다.
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("f3() Done:", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3():", r)
	}
	return
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a delay!")
		return
	}

	delay, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Delay :", delay)

	f1(delay)
	f2(delay)
	f3(delay)
}
