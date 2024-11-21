/*
패킷 크기의 제한, 작업의 식별, 다양한 에러들을 코드화하는 데 사용되는 핵심 유형을 정리한다.
*/

package tftp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"strings"
)

const (
	DatagramSize = 516              // 최대 지원하는 데이터그램 크기
	BlockSize    = DatagramSize - 4 // 4바이트의 헤더 -> 데이터 블록의 최대크기
)

type OpCode uint16 // 헤더의 첫 2바이트 - 작업을 나타내는 OP코드(operation code)이다. op코드는 RFC 1350 문서에 기재된 정수 값에 상응한다.

const (
	OpRRQ  OpCode = iota + 1 // 서버는 읽기 요청 및 아래 3가지를 포함한 4개의 동작을 지원한다.
	_                        // WRQ 미지원 - 쓰기 요청은 지원하지 않는다. 읽기 전용 서버이다.
	OpData                   // 데이터 작업
	OpAck                    // 메시지 승인
	OpErr                    // 에러
)

type ErrCode uint16 // 헤더의 다음 2바이트 - 에러 코드

const (
	ErrUnknown ErrCode = iota
	ErrNotFound
	ErrAccessViolation
	ErrDiskFull
	ErrIllegalOp
	ErrUnknownID
	ErrFileExists
	ErrNoUser
)

/*
=======================  읽기 요청

- 클라이언트가 다운로드 요청을 하면 서버는 읽기 요청 패킷을 수신한다. 서버는 데이터 패킷 혹은 에러 패킷으로 응답해야 된다.
이 두 패킷은 서버가 클라이언트의 읽기 요청 패킷을 수신했다는 것을 확인해 주는 역할을 한다. 만일 클라이언트가 이 두 패킷을 받지 못한다면 재요청 또는 포기해야 된다.
- 읽기 요청 패킷은 2바이트의 op코드와 n바이트의 파일명, 1바이트의 널 문자, 모드 정보와, 다시 널 문자로 구성된다.
- 모드 정보는 서버가 파일을 전송할 때 netascii 모드 혹은 octet 모드로 전송할지에 대한 정보이다. 우리는 octet 모드만을 받도록 한다.
*/

// === 읽기 요청과 바이너리 마샬링 메서드
type ReadReq struct {
	Filename string
	Mode     string
}

// 서버에서 사용되지는 않지만 클라이언트가 이 메서드를 사용한다.
func (q ReadReq) MarshalBinary() ([]byte, error) {
	mode := "octet"
	if q.Mode != "" {
		mode = q.Mode
	}

	// OP코드 + 파일명 + 0바이트 + 모드 정보 + 0바이트
	cap := 2 + 2 + len(q.Filename) + 1 + len(mode) + 1

	// bytes.Buffer -> Go 에서 제공하는 패키지, 바이트 슬라이스를 효율적으로 다룰 수 있도록 도와주는 버퍼이다.
	// 바이트 데이터를 읽고 쓰는데 사용되며, 자동으로 크기를 조절한다.
	// 주요 메서드 :
	/*
		Write(p []byte) (n int, err error) // 버퍼에 바이트 슬라이스 쓰기
		WriteByte(c byte) error				// 버퍼에 단일 바이트 쓰기
		WriteString(s string) (n int, err error)	// 버퍼에 문자열 쓰기
		Read(p []byte)	(n int, err error)		// 버퍼에서 바이트 슬라이스 읽기
		Bytes() []byte							// 버퍼를 바이트 슬라이스로 반환
		String() string							// 버퍼를 문자열로 반환
	*/

	b := new(bytes.Buffer)

	// Grow(int n) - bytes.Buffer의 내부 버퍼 용량을 최소한 n바이트로 확장
	// 이 메서드는 버퍼의 용량을 미리 할당하여, 데이터를 추가할 때 발생하는 여러 번의 메모리 할당을 피해서 성능 향상
	b.Grow(cap)

	// op코드는 단일 바이트가 아닌 구조체나 오려 바이트로 구성된 데이터이기 때문에 binary.Write를 사용한다.
	err := binary.Write(b, binary.BigEndian, OpRRQ) // op코드 쓰기
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(q.Filename)
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(mode)
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// 읽기 요청 구현의 마무리, 서버가 일반적으로 클라이언트와의 네트워크 연결에서 읽어 들이는 바이트 슬라이스의 읽기 요청을 언마샬링 할 수 있는 방법 정의
func (q *ReadReq) UnmarshalBinary(p []byte) error {
	r := bytes.NewBuffer(p) // byte 슬라이스를 bytes.buffer로

	var code OpCode

	// 바이너리 데이터는 구분자가 필요하지 않다. 하지만 데이터 타입을 정확히 알고 있어야 된다.
	// 바이너리를 빅엔디언 방식으로 읽어서 OpCode에 저장한다.
	err := binary.Read(r, binary.BigEndian, &code) // OP코드 읽어서 code에 저장
	if err != nil {
		return errors.New("opcode 에러")
	}

	if code != OpRRQ {
		return errors.New("invalid RRQ")
	}

	q.Filename, err = r.ReadString(0) // 파일명 읽기 (0까지)
	if err != nil {
		return errors.New("invalid RRQ")
	}

	q.Filename = strings.TrimRight(q.Filename, "\x00") // 0바이트 제거
	if len(q.Filename) == 0 {
		return errors.New("invalid RRQ")
	}

	q.Mode, err = r.ReadString(0) // 모드 정보 읽기
	if err != nil {
		return errors.New("mode info Error")
	}
	q.Mode = strings.TrimRight(q.Mode, "\x00") // 0바이트 제거( \x00 아스키 코드의 널문자 0에 해당)
	if len(q.Mode) == 0 {
		return errors.New("invalid RRQ")
	}

	actual := strings.ToLower(q.Mode) // 강제로 octet모드 설정
	if actual != "octet" {
		return errors.New("only binay transfers supported")
	}

	return nil
}

/*
============== 데이터 패킷
- 클라이언트는 읽기 요청에 대한 응답으로 데이터 패킷을 수신한다.
- 서버는 일련의 데이터 패킷으로 파일을 전송한다. 각각의 데이터 패킷은 1을 시작으로 점차 증가하는 숫자의 할당된 블럭 번호를 갖는다.
블럭 번호를 통해서 데이터 정렬 및 중복데이터 처리를 목적으로 사용된다.
- 마지막 패킷을 제외한 모든 데이터 패킷은 512바이트의 페이로드를 갖는다. 클라이언트는 전송의 마지막을 알리는, 512바이트보다 작은 페이로드의
패킷을 수신할 때까지 계속해서 데이터를 읽는다. 클라이언트와 서버는 언제든지 확인 패킷 대신 에러 패킷을 반환할 수 있다. 이경우 데이터 전송을 중단한다.

- 데이터 패킷의 구조 OP코드(2바이트), 블럭 번호(2바이트), 페이로드(n(최대512)바이트)
- 클라이언트는 서버에서 데이터 패킷을 수신할 때마다 확인 패킷을 보내야 한다. 서버가 일정 시간 내에 확인 패킷을 받지 못하거나 에러가 발생한경우,
서버는 재시도 한도 내에서 재전송을 시도한다.
*/

// 실제 데이터 전송에 사용되는 데이터 타입
type Data struct {
	Block   uint16 // 16비트 양의 정수인 블럭 번호가 오버플로(다시 0으로) 될 수 있다. 33.5MB보다 큰 페이로드는 오버플로 된다. 클라이언트에서 우아하게 처리하도록 한다.
	Payload io.Reader
}

// 데이터 마샬링 서버에서
func (d *Data) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	b.Grow(DatagramSize)

	d.Block++ // 블럭번호 1씩 증가

	err := binary.Write(b, binary.BigEndian, OpData) // op코드 쓰기
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, d.Block) // 블럭 번호 쓰기
	if err != nil {
		return nil, err
	}

	// BlockSize 만큼 쓰기
	_, err = io.CopyN(b, d.Payload, BlockSize)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

// 데이터 언마샬 - 클라이언트
func (d *Data) UnmarshalBinary(p []byte) error {
	if l := len(p); l < 4 || l > DatagramSize {
		return errors.New("invalid Data")
	}

	var opcode OpCode

	err := binary.Read(bytes.NewBuffer(p[:2]), binary.BigEndian, &opcode) // 2바이트의 해더만 읽기(opcode)
	if err != nil || opcode != OpData {
		return errors.New("invalid Data")
	}

	err = binary.Read(bytes.NewBuffer(p[2:4]), binary.BigEndian, &d.Block) // 블럭의 순서 읽어오기
	if err != nil {
		return errors.New("invalid Data")
	}

	d.Payload = bytes.NewBuffer(p[4:])

	return nil
}

/*
================= 수신 확인
- 수신 확인 패킷은 4바이트의 길이를 갖는다. opcode(2바이트), 블럭번호(2바이트)
*/
// 수신 확인 타입에 대한 전체 구현
type Ack uint16 // 수신 확인 패킷 -> 이 정수는 수신 확인된 블록 번호를 나타낸다.

func (a Ack) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	cap := 4

	b.Grow(cap)

	err := binary.Write(b, binary.BigEndian, OpAck)
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, a)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (a Ack) UnMarshalBinary(p []byte) error {
	b := bytes.NewReader(p)

	var opcode OpCode
	err := binary.Read(b, binary.BigEndian, &opcode)
	if err != nil {
		return err
	}

	if opcode != OpAck {
		return errors.New("invalid Ack")
	}

	return binary.Read(b, binary.BigEndian, a) // 블럭 번호 읽기
}

/*
========================== 에러 처리
- 에러 패킷 : opcode(2바이트), 에러코드(2), 메시지(n바이트), 0(널문자 1바이트)
*/

// ==== 클라이언트와 서버 간의 에러를 전달하기 위해 사용되는 에러 타입
type Err struct {
	Error   ErrCode
	Message string
}

func (e Err) MarshalBinary() ([]byte, error) {
	b := new(bytes.Buffer)
	cap := 2 + 2 + len(e.Message) + 1
	b.Grow(cap)

	err := binary.Write(b, binary.BigEndian, OpErr)
	if err != nil {
		return nil, err
	}

	err = binary.Write(b, binary.BigEndian, e.Error)
	if err != nil {
		return nil, err
	}

	_, err = b.WriteString(e.Message)
	if err != nil {
		return nil, err
	}

	err = b.WriteByte(0)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (e *Err) UnMarshalBinary(p []byte) error { // 값을 변경해야 되기 때문에 포인터로 가져온다.
	b := bytes.NewBuffer(p)

	var opcode OpCode

	err := binary.Read(b, binary.BigEndian, &opcode)
	if err != nil {
		return err
	}
	if opcode != OpErr {
		return errors.New("invalid Err")
	}

	err = binary.Read(b, binary.BigEndian, e.Error)
	if err != nil {
		return err
	}

	e.Message, err = b.ReadString(0)
	e.Message = strings.TrimRight(e.Message, "\x00")
	return err
}
