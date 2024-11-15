/*
bufio.Scanner 는 간편하게 구분자로 구분된 데이터를 읽어 들일 수 있는 Go의 표준 라이브러리이다. Scanner는 매개변수로 io.Reader를 받는다.
net.Conn 인터페이스에도 io.Reader 인터페이스를 구현한 Read 메서드를 가지고 있기 때문에 Scanner를 이용하여 쉽게 네트워크 연결로부터 구분자로 구분된 데이터를 받을 수 있다.
*/

package main

import (
	"bufio"
	"net"
	"reflect"
	"testing"
)

const payload = "The bigger the interface, the weaker the abstraction."

func TestScanner(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		bytes := []byte(payload)
		_, err = conn.Write(bytes)
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)

	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		t.Error(err)
	}
	expected := []string{"The", "bigger", "the", "interface,", "the", "weaker", "the", "abstraction."}
	if !reflect.DeepEqual(words, expected) {
		t.Fatal("다르다.")
	}
	t.Logf("Scanned words: %#v", words)
}
