// 콘텍스트를 이용하는 또 다른 장점으로는 cancel 함수 자체에 있다.
// 데드라인을 지정하지 않고도 필요 시에  cancel 함수를 이용하여 연결 시도를 취소할 수 있다.

package ch03

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestDialContextCancel(t *testing.T) {
	// 연결 시도를 중단하기 위해 데드라인을 설정해서 콘텍스트를 생성하고 데드라인이 지나기까지 기다리는 대신
	// context.WithCancel 함수를 이용하여 콘텍스트(ctx)와 콘텍스트를 취소할 수 있는 함수(cancel)를 받는다.
	ctx, cancel := context.WithCancel(context.Background())
	sync := make(chan struct{})

	go func() {
		defer func() { sync <- struct{}{} }()
		var d net.Dialer
		d.Control = func(_, _ string, _ syscall.RawConn) error {
			time.Sleep(time.Second)
			return nil
		}

		conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:http")
		if err != nil {
			t.Log(err)
			return
		}

		conn.Close()
		t.Error("connection did not timeout")
	}()

	cancel()
	<-sync

	if ctx.Err() != context.Canceled {
		t.Errorf("expected canceled context; actual : %q", ctx.Err())
	}
}
