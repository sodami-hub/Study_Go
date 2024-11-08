/*
	ch03/ping.go + ch03/ping_example_test.go 를 실행 코드로 합친 코드이다.

네트워크 연결을 지속하기 위해서 애플리케이션 계층에서 긴 유휴 시간을 가져야만 하는경우가 있다.
이러한 경우 지속시간을 뒤로 설정하기 위해 노드 간에 하트비트를 구현해야 된다.

먼저 일정한 간격으로 핑을 전송하기 위한 고루틴을 실행할 수 있는 약간의 코드가 필요하다.
그리고 최근에 데이터를 받은 원격 노드로 불필요하게 또 다시 핑을 할 필요가 없을 테니, 핑 타이머를 초기화할 방법이 필요하다.
*/

package main

import (
	"context"
	"fmt"
	"io"
	"time"
)

const defaultPingInterval = 30 * time.Second

func Pinger(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var interval time.Duration
	select {
	case <-ctx.Done():
		return
		//1.
	case interval = <-reset: // reset 채널에서 초기 간격을 받아옴
	default:
	}

	if interval <= 0 {
		interval = defaultPingInterval
	}

	//2. 타이머를 interval로 초기화
	timer := time.NewTimer(interval)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for { // 여러 채널에서 대기!! 이때 recievePing 함수에서 resetTimer 채널을 통해 새로운 간격을 받아오면 그순간 돌아가기 시작!!
		select {
		case <-ctx.Done(): // 컨텍스트 종료
			return
		case newInterval := <-reset: // 타이머 리셋을 위한 새로운 간격을 받아옴
			if !timer.Stop() {
				<-timer.C
			}
			if newInterval > 0 {
				interval = newInterval
			}
		case <-timer.C: // 타이머가 만료되면 핑을 보냄
			if _, err := w.Write([]byte("ping")); err != nil {
				// 여기서 연속적으로 발생하는 타임아웃을 추적하고 처리
				return
			}
		}
		_ = timer.Reset(interval) // 핑을 보내고 타이머 리셋
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	r, w := io.Pipe() // io.Pipe 함수는 io.Reader와 io.Writer 인터페이스를 구현하는 두 객체를 반환한다.
	// 이 두 객체는 메모리 버퍼를 사용해서 데이터를 읽고 쓸 수 있다.
	done := make(chan struct{})
	resetTimer := make(chan time.Duration, 1)
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

		fmt.Printf("received %q (%s)\n", buf[:n], time.Since(now).Round(time.Millisecond))
	}

	for i, v := range []int64{0, 200, 300, 0, -1, -1, -1} {
		fmt.Printf("Run %d:\n", i+1)
		receivePing(time.Duration(v)*time.Millisecond, r)
	}

	cancel()
	<-done
}
