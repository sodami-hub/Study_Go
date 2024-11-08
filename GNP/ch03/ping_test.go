// 하트비트를 이용하여 데드라인 늦추기

package ch03

import (
	"context"
	"io"
	"net"
	"testing"
	"time"
)

func TestPingerAdvanceDeadline(t *testing.T) {
	done := make(chan struct{})
	listner, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	begin := time.Now()
	go func() {
		defer func() {
			close(done) // 채널에 done 전달
		}()

		conn, err := listner.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		ctx, cancel := context.WithCancel(context.Background())

		defer func() {
			cancel()
			conn.Close()
		}()

		resetTimer := make(chan time.Duration, 1)
		resetTimer <- time.Second
		go Pinger(ctx, conn, resetTimer)

		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			t.Error(err)
			return
		}

		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			// nErr, _ := err.(net.Error)
			// if nErr.Timeout() { // 타임아웃이 발생하면 nErr.Timeout()이 true가 된다. 여기서는 테스트의 성공을 위해서 !를 붙여준다.
			// 	t.Errorf("expected timeout error; actual : %v", err)
			// 	return
			// }
			if err != nil {
				return
			}
			t.Logf("[%s] %s", time.Since(begin).Truncate(time.Second), buf[:n])
			resetTimer <- 0
			err = conn.SetDeadline(time.Now().Add(5 * time.Second))
			if err != nil {
				t.Error(err)
				return
			}
		}
	}()

	conn, err := net.Dial("tcp", listner.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	buf := make([]byte, 1024)
	for i := 0; i < 4; i++ { // 핑을 4개 읽음
		n, err := conn.Read(buf)
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("[%s] %s", time.Since(begin).Truncate(time.Second), buf[:n])
	}
	// time.Sleep(3 * time.Second) 여기에 이 코드가 있으면
	_, err = conn.Write([]byte("Pong!!!")) // 핑 타이머 초기화해야 됨
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 4; i++ { // 핑을 4개 더 읽음
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
			break
		}
		t.Logf("[%s] %s", time.Since(begin).Truncate(time.Second), buf[:n])
	}
	<-done // 비동기적으로 작동됨. 채널에 값이 들어오면 이것에 따른 코드가 실행됨.
	end := time.Since(begin).Truncate(time.Second)
	t.Logf("[%s] done", end)
	if end != 9*time.Second {
		t.Fatalf("expected EOF at 9 seconds; actual %s", end)
	}
}
