// 연결 시도를 타임아웃하기 위해 데드라인 콘텍스트 사용하기

package ch03

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func TestDialContext(t *testing.T) {
	dl := time.Now().Add(5 * time.Second)                         //연결을 시도하기 전에, 5초 후 데드라인이 지나는 콘텍스트를 만들기 위해 현재 시간으로부터 5초 뒤의 시간을 저장
	ctx, cancel := context.WithDeadline(context.Background(), dl) // WithDeadline 함수를 이용해서 콘텍스트와 cancel 함수를 생성하고 위에서 생성한 데드라인을 설정
	defer cancel()

	var d net.Dialer                                         // DialContext는 Dialer의 메서드이다.
	d.Control = func(_, _ string, _ syscall.RawConn) error { // Dialer의 Control함수를 오버라이딩하여 연결을 콘텍스트의 데드라인을 간신히 초과하는 정도로 지연시킨다.
		// 콘텍스트의 데드라인이 지나기 위해 충분히 긴 시간 동안 대기한다.
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}
	conn, err := d.DialContext(ctx, "tcp", "10.0.0.0:80") // 위에서 생성한 Dialer의 DialContext를 사용해서 ctx를 등록해서 연결을 시도.
	if err == nil {
		conn.Close()
		t.Fatal("connection did not time out")
	}
	nErr, ok := err.(net.Error)
	if !ok {
		t.Error(err)
	} else {
		if !nErr.Timeout() {
			t.Errorf("error is not a timeout: %v", err)
		}
	}
	if ctx.Err() != context.DeadlineExceeded {
		t.Errorf("expected deadline exceeded; actual : %v", ctx.Err())
	}
}
