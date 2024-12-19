/*
데이터그램 소켓 - unixgram
udp나 unixgram과 같은 데이터그램 기반 네트워크 타입과 통신하는 에코 서버
이 서버는 UDP로 통신하든 unixgram 소켓으로 통신하든 결국 동일하게 보일 것이다.
차이점이 있따면 unixgram 리스너와 통신할 경우 통신이 종료될 때 소켓 파일을 제거해야 된다는 것이다.
*/

package datagramsock

import (
	"context"
	"net"
	"os"
)

func datagramEchoServer(ctx context.Context, network, addr string) (net.Addr, error) {

	s, err := net.ListenPacket(network, addr)
	if err != nil {
		return nil, err
	}

	go func() {
		defer func() {
			<-ctx.Done()
			_ = s.Close()
			if network == "unixgram" {
				_ = os.Remove(addr)
			}
		}()

		buf := make([]byte, 1024)
		for {
			n, clientAddr, err := s.ReadFrom(buf)
			if err != nil {
				return
			}

			_, err = s.WriteTo(buf[:n], clientAddr)
			if err != nil {
				return
			}
		}
	}()
	return s.LocalAddr(), nil

}
