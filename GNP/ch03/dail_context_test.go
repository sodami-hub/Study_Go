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
	dl := time.Now().Add(5 * time.Second) //연결을 시도하기 전에, 5초 후 데드라인이 지나는 콘텍스트를 만들기 위해 현재 시간으로부터 5초 뒤의 시간을 저장
	// WithDeadline 함수를 이용해서 콘텍스트와 cancel 함수를 생성하고 위에서 생성한 데드라인을 설정
	// 주어진 Deadline까지 작업을 완료하지 못하면 자동으로 취소되도록 설정하는 context
	ctx, cancel := context.WithDeadline(context.Background(), dl)
	defer cancel()

	var d net.Dialer // DialContext는 Dialer의 메서드이다.
	// Dialer의 Control함수를 오버라이딩하여 연결을 콘텍스트의 데드라인을 간신히 초과하는 정도로 지연시킨다.
	d.Control = func(_, _ string, _ syscall.RawConn) error {
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

/*
context.Background():

context.Background()는 기본 컨텍스트를 생성합니다. 이 컨텍스트는 최상위 부모 컨텍스트로, 주로 다른 컨텍스트를 생성할 때 사용됩니다.
context.WithDeadline:

context.WithDeadline 함수는 주어진 데드라인(마감 시간)까지 유효한 컨텍스트를 생성합니다.
이 함수는 두 개의 값을 반환합니다: 새로운 컨텍스트(ctx)와 취소 함수(cancel).
dl:

dl은 데드라인을 나타내는 time.Time 값입니다. 이 시점까지 작업이 완료되지 않으면 컨텍스트가 자동으로 취소됩니다.
ctx:

ctx는 새로운 컨텍스트로, 주어진 데드라인까지 유효합니다. 이 컨텍스트를 사용하여 작업을 수행할 수 있습니다.
cancel:

cancel은 취소 함수로, 이 함수를 호출하면 컨텍스트가 즉시 취소됩니다. 작업이 완료되면 반드시 이 함수를 호출하여 리소스를 해제해야 합니다.
*/
