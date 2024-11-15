package dynamicbuff

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

//===========================  간단한 프로토콜을 구현하는 메시지 구조체 ==============================

// 정의할 타입을 나타내는 상수 1,2.
// 보안상의 문제로 3. 최대 페이로드 크기를 반드시 정의해 줘야 된다.
const (
	BinaryType     uint8  = iota + 1 // 1
	StringType                       // 2
	MaxPayloadSize uint32 = 10 << 20 // 10MB  3.
)

var ErrMaxPayloadSize = errors.New("maximum payload size exceeded")

// 각 타입별 메시지들이 구현해야 하는 Payload라는 이름의 인터페이스를 정의한다.
// 각 메서드를 반드시 구현해야 된다.
type Payload interface {
	Bytes() []byte
	fmt.Stringer // String() string {}
	// 각 타입별 메시지를 reader로 부터 읽고, wrter에 쓸 수 있게 해 주는 기능의 형태를 제공
	io.ReaderFrom
	io.WriterTo
}

//=========================== TLV 기반 타입의 첫번째 :  Binary 타입의 Payload 인터페이스 구현 ==============================

type Binary []byte // Binary 타입은 바이트 슬라이스

func (m Binary) Bytes() []byte { // 자기 자신 반환
	return m
}
func (m Binary) String() string { // 자기 자신을 문자열로 캐스팅해서 반환
	return string(m)
}

// WriteTo 메서드는 io.Writer인터페이스를 메개변수로 받아서 writer에 쓰인 바이트 수와 에러 인터페이스를 반환한다.
func (m Binary) WriteTo(w io.Writer) (int64, error) {

	// 헤더 부분
	err := binary.Write(w, binary.BigEndian, BinaryType) // 1바이트 타입을 writer(w)에 쓴다. => 데이터 타입
	if err != nil {
		return 0, err
	}
	var n int64 = 1

	err = binary.Write(w, binary.BigEndian, uint32(len(m))) // 4바이트 크기를 writer(w)에 쓴다 => 갑의 크기를 나타내는 값
	if err != nil {
		return n, err
	}
	n += 4

	// 페이로드 - 실제 전송할 데이터
	o, err := w.Write(m) // 페이로드

	return n + int64(o), err
}

func (m *Binary) ReadFrom(r io.Reader) (int64, error) {

	// 헤더 부분 읽어오고
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ) // 1바이트 타입 - 바이너리 타입인지 확인
	if err != nil {
		return 0, err
	}
	var n int64 = 1
	if typ != BinaryType {
		return n, errors.New("invalid Binary")
	}

	var size uint32
	err = binary.Read(r, binary.BigEndian, &size) // 4바이트 크기 - 페이로드의 크기 확인
	if err != nil {
		return n, err
	}
	n += 4
	// 실제 4바이에 저장될 수 있는 최댓값은 4GB 정도인데 이 크기를 다 사용하면 악용될 수 있다.
	if size > MaxPayloadSize {
		return n, ErrMaxPayloadSize
	}

	// 헤더에 문제가 없으면 페이로드 읽어오기
	*m = make([]byte, size) // 헤더에서 읽어드린 페이로드의 크기만큼 버퍼 생성
	o, err := r.Read(*m)

	return n + int64(o), err // 읽어온 총 바이트수 리턴
}

//=========================== TLV 기반 타입의 두 번째 :  문자열 타입의 Payload 인터페이스 구현 ==============================

type String string

func (m String) Bytes() []byte {
	return []byte(m)
}

func (m String) String() string {
	return string(m)
}

func (m String) WriteTo(w io.Writer) (int64, error) {
	// 헤더 부분
	err := binary.Write(w, binary.BigEndian, StringType) // 1바이트 타입을 writer(w)에 쓴다. => 데이터 타입
	if err != nil {
		return 0, err
	}
	var n int64 = 1

	err = binary.Write(w, binary.BigEndian, uint32(len(m))) // 4바이트 크기를 writer(w)에 쓴다 => 갑의 크기를 나타내는 값
	if err != nil {
		return n, err
	}
	n += 4

	// 페이로드 - 실제 전송할 데이터
	o, err := w.Write([]byte(m)) // 페이로드

	return n + int64(o), err
}

func (m *String) ReadFrom(r io.Reader) (int64, error) {
	// 헤더 부분 읽어오고
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ) // 1바이트 타입 - 바이너리 타입인지 확인
	if err != nil {
		return 0, err
	}
	var n int64 = 1
	if typ != StringType {
		return n, errors.New("invalid Binary")
	}

	var size uint32
	err = binary.Read(r, binary.BigEndian, &size) // 4바이트 크기 - 페이로드의 크기 확인
	if err != nil {
		return n, err
	}
	n += 4
	// 실제 4바이에 저장될 수 있는 최댓값은 4GB 정도인데 이 크기를 다 사용하면 악용될 수 있다.
	if size > MaxPayloadSize {
		return n, ErrMaxPayloadSize
	}

	// 헤더에 문제가 없으면 페이로드 읽어오기

	buf := make([]byte, size) // 헤더에서 읽어드린 페이로드의 크기만큼 버퍼 생성
	o, err := r.Read(buf)     // 페이로드 읽어오기
	if err != nil {
		return n, err
	}

	*m = String(buf)

	return n + int64(o), nil // 읽어온 총 바이트수 리턴
}

// ======== reader에서 바이트를 읽어서 Binary와 String 타입으로 디코딩하기

func decode(r io.Reader) (Payload, error) {
	var typ uint8
	err := binary.Read(r, binary.BigEndian, &typ)
	if err != nil {
		return nil, err
	}

	var payload Payload

	switch typ {
	case BinaryType:
		payload = new(Binary)
	case StringType:
		payload = new(String)
	default:
		return nil, errors.New("unknown type")
	}
	// 위에서 타입을 추론하기 위해서 1바이트를 읽었다. 그렇기 때문에 그대로 ReadFrom에 보낼 수 없다. ReadFrom에서도 첫 1바이트를 읽어야 되기 때문이다.
	// 이 상황에서 io.MultiReader 라는 함수를 사용할 수 있다. 이 함수는 reader로부터 이미 읽은 바이트([]byte{typ})를 다음에 읽을 바이트와 연결하는 데 사용한다.
	// 다시말해서 읽었던 바이트를 다시 reader로 주입한다.
	_, err = payload.ReadFrom(io.MultiReader(bytes.NewReader([]byte{typ}), r))
	if err != nil {
		return nil, err
	}

	return payload, nil
}
