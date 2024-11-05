/*
- 다중 다이얼러 취소
여러 개의 서버에서 TCP를 통해 단 하나의 리소스만 받아 올 필요가 있다고 가정해보자.
각 서버에 비동기적으로 연결을 요청하고, 동일한 콘텍스트를 각 다이얼러에 전달한다.
한 서버로부터 응답이 왔으면 다른 응답은 필요 없으니 나머지 다이얼러들은 콘텍스트를 취소하여 연결 시도를 중단할 수 있다.

아래 코드는 동일한 콘텍스트를 여러 개의 다이얼러들에게 전달한다. 첫 응답을 받으면 콘텍스트를 취소ㅗ하여 그 외 나머지 다이얼러들의 연결 시도를 중단한다.
*/

package ch03

import (
	"context"
	"net"
	"sync"
	"testing"
	"time"
)

func TestDialContextCancelFanOut(t *testing.T) {
	ctx, cancel := context.WithDeadline(
		context.Background(),
		time.Now().Add(5*time.Second),
	)

	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		//하나의 연결만 수락한다. for loop가 아님.
		conn, err := listener.Accept()
		if err != nil {
			conn.Close()
		}
	}()

	dial := func(ctx context.Context, address string,
		response chan int, id int, wg *sync.WaitGroup) { // tcp 연결을 시도하는 함수
		defer wg.Done()

		var d net.Dialer
		c, err := d.DialContext(ctx, "tcp", address)
		if err != nil {
			return
		}
		c.Close()

		select { // select 문은 두 개의 채널 연산 중 하나가 준비될 때까지 대기한다.
		case <-ctx.Done(): // 채널이 닫히면 첫 번째 케이스가 실행된다.
		case response <- id: // 채널이 값을 받을 준비가 되면 두 분째 케이스가 실행된다.
		}
	}

	res := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go dial(ctx, listener.Addr().String(), res, i+1, &wg)
	}

	response := <-res // 가장 먼저 연결에 성공한 ID를 res 채널을 통해서 받는다.
	cancel()          // 컨텍스트를 취소하여 다른 고루틴이 더 이상 연결을 시도하지 않도록 한다.
	wg.Wait()
	close(res)

	if ctx.Err() != context.Canceled { // 컨텍스트가 취소되었는지 확인한다.
		t.Errorf("expected canceled context; actual: %s", ctx.Err())
	}
	t.Logf("dialer %d retrieved the resource", response) // 성공한 연결의 ID를 로그한다.
}
