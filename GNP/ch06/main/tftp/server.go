/*
TFTP 서버
서버는 네트워크 연결로 정의한 타입이 마샬링된 바이트 슬라이스를 전송하기만 하면 된다. 타입 인터페이스의 구현체가 그 외의 부분을 처리할 것이다.
*/

package tftp

import (
	"bytes"
	"errors"
	"log"
	"net"
	"time"
)

type Server struct {
	Payload []byte        // 모든 읽기 요청에 반환될 페이로드
	Retries uint8         // 전송 실패 시 재시도 횟수
	Timeout time.Duration // 전송 승인을 기다릴 기간
}

func (s Server) ListenAndServe(addr string) error {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	log.Printf("Listening on %s ... \n", conn.LocalAddr())

	return s.Serve(conn)
}

func (s *Server) Serve(conn net.PacketConn) error {
	if conn == nil {
		return errors.New("nil connection")
	}
	if s.Payload == nil {
		return errors.New("payload is required")
	}
	if s.Retries == 0 {
		s.Retries = 10
	}
	if s.Timeout == 0 {
		s.Timeout = 5 * time.Second
	}

	// 클라이언트의 요청에 대한 언마샬... 에러가 아니면 핸들러로 전달
	var rrq ReadReq

	for {
		buf := make([]byte, DatagramSize)

		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			return err
		}

		err = rrq.UnmarshalBinary(buf) // rrq에 filename, mode의 정보가 반환될 것이다.
		if err != nil {
			log.Printf("[%s] bad request: %v", addr, err)
			continue
		}

		go s.handle(addr.String(), rrq) // 네트워크 연결에서 읽은 데이터가 읽기 요청인 경우 서버는 데이터를 고루틴의 핸들러로 전달한다.
	}
}

// 클라이언트로부터의 읽기 요청을 수락하고 서버의 페이로드로 응답한다.
// 핸들러는 하나의 데이터 패킷을 전송하고 다음 데이터 패킷을 보내기 전에 클라이언트로부터 수신 확인 패킷을 대기한다. 또한, 일정 시간 안에 클라이언트로부터
// 수신확인 패킷 응답을 받지 못하면 현재의 데이터 패킷을 다시 보내려 할 것이다.
func (s Server) handle(clientAddr string, rrq ReadReq) {
	log.Printf("[%s] requested file: %s", clientAddr, rrq.Filename)

	conn, err := net.Dial("udp", clientAddr)
	if err != nil {
		log.Printf("[%s] dial: %v", clientAddr, err)
		return
	}
	defer func() { _ = conn.Close() }()

	var (
		ackPkt  Ack
		errPkt  Err
		dataPkt = Data{Payload: bytes.NewReader(s.Payload)} // 서버의 페이로드를 사용해서 data 객체를 준비
		buf     = make([]byte, DatagramSize)
	)

NEXTPACKET:
	for n := DatagramSize; n == DatagramSize; { // 초기값으로 n을 데이타그램사이즈로 초기화하고 for 돌림 - 이후에도 n의 값이 데이타그램사이즈면 다시 반복
		data, err := dataPkt.MarshalBinary()
		if err != nil {
			log.Printf("[%s] preparing data packet : %v", clientAddr, err)
			return
		}
	RETRY:
		for i := s.Retries; i > 0; i-- {
			n, err = conn.Write(data) // 데이터 패킷 전송
			if err != nil {
				log.Printf("[%s] write: %v", clientAddr, err)
				return
			}

			// 클라이언트의 ACK 패킷 대기
			_ = conn.SetDeadline(time.Now().Add(s.Timeout))
			_, err = conn.Read(buf)

			if err != nil {
				if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
					continue RETRY
				}

				log.Printf("[%s] waiting for ACK: %v", clientAddr, err)
				return
			}

			// ackErr := ackPkt.UnmarshalBinary(buf)
			// errErr := errPkt.UnmarshalBinary(buf)

			// if ackErr == nil {
			// 	if uint16(ackPkt) == dataPkt.Block {
			// 		continue NEXTPACKET
			// 	}
			// } else if errErr == nil {
			// 	log.Printf("[%s] received error : %v", clientAddr, errPkt.Message)
			// 	return
			// } else {
			// 	log.Printf("[%s] bad packet!!! %s, %s", clientAddr, ackErr, errErr)
			// }

			// 클라이언트로 받은 바이트를 Ack, Err 객체로 언마샬리을 시도한다.
			switch {
			case ackPkt.UnmarshalBinary(buf) == nil: // 클라이언트로부터 ACK 패킷이 왔다.
				if uint16(ackPkt) == dataPkt.Block { // 블럭번호를 확인하고 방금 보낸 블럭번호가 맞으면 다음 패킷으로
					// 블럭 번호가 맞지않으면 for 루프를 한번 더 돌아서 현재 패킷을 다시 보냄
					continue NEXTPACKET
				}
			case errPkt.UnmarshalBinary(buf) == nil: // 클라이언트로 에러 패킷이 왔다.
				log.Printf("[%s] received error : %v", clientAddr, errPkt.Message)
				return
			default:
				log.Printf("[%s] bad packet!!!", clientAddr)
			}
		}
		log.Printf("[%s] exhausted retries", clientAddr)
		return
	}
	log.Printf("[%s] sent %d blocks", clientAddr, dataPkt.Block)
}
