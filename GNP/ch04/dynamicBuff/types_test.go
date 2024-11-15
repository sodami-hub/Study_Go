/*
test형태로 types.go의 decode 함수의 동작 살펴보기
두 종류의 다른 타입의 데이터를 네트워크 연결을 통해 송신하는 방법과 수신 측에서 원래의 데이터 타입으로 올바르게 디코딩하는 방법을 살펴보겠다.
*/

package dynamicbuff

import (
	"bytes"
	"encoding/binary"
	"net"
	"reflect"
	"testing"
)

func TestPayloads(t *testing.T) {
	b1 := Binary("Clear is better than clever.")
	b2 := Binary("Dont't panic.")

	s1 := String("Errors are values.")
	payloads := []Payload{&b1, &s1, &b2}

	listner, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Error(err)
	}

	go func() {
		conn, err := listner.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer conn.Close()

		for _, p := range payloads {
			_, err = p.WriteTo(conn)
			if err != nil {
				t.Error(err)
				break
			}
		}
	}()

	conn, err := net.Dial("tcp", listner.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	for i := 0; i < len(payloads); i++ {
		actual, err := decode(conn)
		if err != nil {
			t.Fatal(err)
		}

		if expected := payloads[i]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("value mismatch: %v != %v", expected, actual)
			continue
		}

		t.Logf("[%T] %[1]q", actual)
	}

}

// 최대 페이로드 크기 테스트
func TestMaxPayloadSize(t *testing.T) {
	buf := new(bytes.Buffer)
	err := buf.WriteByte(BinaryType)
	if err != nil {
		t.Fatal(err)
	}

	err = binary.Write(buf, binary.BigEndian, uint32(1<<30)) // 1GB
	if err != nil {
		t.Fatal(err)
	}

	var b Binary
	_, err = b.ReadFrom(buf)
	if err != ErrMaxPayloadSize { // 실제 이 코드는 ErrMaxPayloadSize 에러가 발생한다. == 으로 바꾸면 에러 캐치
		t.Fatalf("expected ErrMaxPayloadSze; actual: %v", err)
	}
}
