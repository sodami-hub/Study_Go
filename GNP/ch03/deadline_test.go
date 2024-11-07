/*
Go의 네트워크 연결 객체는 읽기와 쓰기 동작에 대해 모두 데드라인을 포함한다.
데드라인은 아무런 패킷도 오고가지 않은 채로 네트워크 연결이 얼마나 유휴 상태로 지속할 수 있는지를 제어한다.

Read 메서드에 대한 데드라인은 연결 객체 내의 SetReadDeadline을
Write 메서드에 대해서는 SetWriteDeadline을 사용해서 제어하고,
SetDeadline 메서드를 사용해서 동시에 제어할 수 있다.

네트워크 연결상의 데드라인이 지나면 곧바로 타임아웃 에러를 발생한다.

Go 네트워크 연결은 기본적으로 읽기,쓰기에 대해 데드라인을 설정하지 않는다. 즉, 네트워크 연결이 끊기지 않고 오랫동안 존재할 수 있다.
이는 종종 케이블이 뽑히는 등의 장애 상황을 감지할 수 없게 된다.

다음 코드는 서버에서의 연결 객체에 대드라인을 구현한다.
*/

package ch03

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestDeadline(t *testing.T) {
	sync := make(chan struct{})

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer func() {
			conn.Close()
			close(sync) // 빠른 return으로 인해서 sync 채널에서 읽는 데이터가 블로킹되면 안됨
		}()

		//1.
		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			t.Error(err)
			return
		}

		buf := make([]byte, 1)
		_, err = conn.Read(buf) // 원격 노드가 데이터를 보낼 때까지 블로킹됨
		nErr, ok := err.(net.Error)
		//2.
		if !ok || !nErr.Timeout() { // 타임아웃이 발생하면 nErr.Timeout()이 true가 된다. 여기서는 테스트의 성공을 위해서 !를 붙여준다.
			t.Errorf("expected timeout error; actual : %v", err)
		}

		sync <- struct{}{}

		//3.
		err = conn.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			t.Error(err)
			return
		}

		_, err = conn.Read(buf)
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	<-sync
	_, err = conn.Write([]byte("1"))
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1)
	_, err = conn.Read(buf) // 원격 노드가 데이터를 보낼 때까지 블로킹됨
	//4.
	if err != io.EOF { // io.EOF는 원격 노드가 연결을 닫았을 때 발생한다. 실제 이 코드는 io.EOF를 발생하기 때문에 테스트가 실패하지만 !를 붙여서 성공하도록 했다.
		t.Errorf("expected server termination; actual: %v", err)
	}
}
