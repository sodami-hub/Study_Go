package ch03

import (
	"context"
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

	//2.
	timer := time.NewTimer(interval)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for {
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
			if _, err := w.Write([]byte("pingpingping~")); err != nil {
				// 여기서 연속적으로 발생하는 타임아웃을 추적하고 처리
				return
			}
		}
		_ = timer.Reset(interval) // 핑을 보내고 타이머 리셋
	}
}
