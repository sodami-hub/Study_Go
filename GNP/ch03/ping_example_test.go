package ch03

import (
	"context"
	"fmt"
	"io"
	"time"
)

func ExamplePinger() {
	ctx, cancel := context.WithCancel(context.Background())
	r, w := io.Pipe() // io.Pipe 함수는 io.Reader와 io.Writer 인터페이스를 구현하는 두 객체를 반환한다.
	// 이 두 객체는 메모리 버퍼를 사용해서 데이터를 읽고 쓸 수 있다.
	done := make(chan struct{})
	resetTimer := make(chan time.Duration, 1)
	// 버퍼링 된 채널을 생성 (버퍼 크기 1) -> 버퍼링된 채널은 채널에 데이터를 보내면 송신자는 블로킹되지 않고 다른 작업을 수행할 수 있다.
	// 수신자는 버퍼에 값이 있을 때 즉시 값을 받을 수 있다.
	/*
	   비버퍼링된 채널 사용 시 (resetTimer := make(chan time.Duration)):
	   송신자가 값을 보낼 때 수신자가 즉시 값을 받을 준비가 되어 있지 않으면 송신자는 블로킹됩니다.
	   수신자가 값을 받을 때 송신자가 즉시 값을 보낼 준비가 되어 있지 않으면 수신자는 블로킹됩니다.
	   이로 인해 프로그램이 예상치 못하게 블로킹될 수 있습니다.
	*/
	resetTimer <- time.Second // 초기 핑 간격

	go func() {
		Pinger(ctx, w, resetTimer)
		close(done)
	}()

	receivePing := func(d time.Duration, r io.Reader) {
		if d >= 0 {
			fmt.Printf("resetting timer (%s)\n", d)
			resetTimer <- d
		}

		now := time.Now()
		buf := make([]byte, 1024)
		n, err := r.Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("received %q (%s)\n",
			buf[:n], time.Since(now).Round(100*time.Millisecond))
	}

	for i, v := range []int64{0, 200, 300, 0, -1, -1, -1} {
		fmt.Printf("Run %d:\n", i+1)
		receivePing(time.Duration(v)*time.Millisecond, r)
	}

	cancel()
	<-done // ensures the pinger exits after canceling the context

	// Output:
	// Run 1:
	// resetting timer (0s)
	// received "ping" (1s)
	// Run 2:
	// resetting timer (200ms)
	// received "ping" (200ms)
	// Run 3:
	// resetting timer (300ms)
	// received "ping" (300ms)
	// Run 4:
	// resetting timer (0s)
	// received "ping" (300ms)
	// Run 5:
	// received "ping" (300ms)
	// Run 6:
	// received "ping" (300ms)
	// Run 7:
	// received "ping" (300ms)

}
