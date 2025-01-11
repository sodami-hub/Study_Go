package tls_echo

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"
)

// ------------- Server 구조체와 생성자 함수
func NewTLSServer(ctx context.Context, address string, maxIdle time.Duration, tlsConfig *tls.Config) *Server {
	return &Server{
		ctx:       ctx,
		ready:     make(chan struct{}), // make() 를 사용해서 생성하기 때문에 초기 생성시 nil이 아니다.
		addr:      address,
		maxIdle:   maxIdle,
		tlsConfig: tlsConfig,
	}
}

type Server struct {
	ctx       context.Context
	ready     chan struct{}
	addr      string
	maxIdle   time.Duration
	tlsConfig *tls.Config
}

func (s *Server) Ready() {
	log.Println("into the Ready()")
	if s.ready != nil {
		<-s.ready
	}
}

// ------------- 리스닝과 서빙을 처리하고 서버 연결의 읽기 가능한 상태를 알리는 메서드 추가
func (s *Server) ListenAndServeTLS(certFn, keyFn string) error {

	log.Println("into the ListenAndServeTLS()")
	if s.addr == "" {
		s.addr = "localhost:443"
	}

	l, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("binding to tcp %s: %w", s.addr, err)
	}

	if s.ctx != nil {
		go func() {
			<-s.ctx.Done()
			_ = l.Close()
		}()
	}
	return s.ServeTLS(l, certFn, keyFn)
}

// ---------------- net.Listener에 TLS 지원 추가하기

/*
ServeTLS 메서드는 서버의 TLS 구성을 확인한다. 구성이 nil 이면 PreferServerCipherSuites 필드 값을 true로 설정하여 기본 구성을 사용한다.
해당 필드는 서버에서 사용하며, 클라이언트가 원하는 암호와 스위트를 기다리지 ㅇ낳고 서버에서 먼저 TLS 협상 단계에서 사용할 암호화 스위트를 사용한다.
*/
func (s Server) ServeTLS(l net.Listener, certFn, keyFn string) error {
	log.Println("into the ServeTLS()")

	if s.tlsConfig == nil {
		s.tlsConfig = &tls.Config{
			CurvePreferences:         []tls.CurveID{tls.CurveP256},
			MinVersion:               tls.VersionTLS12,
			PreferServerCipherSuites: true,
		}
	}

	if len(s.tlsConfig.Certificates) == 0 && s.tlsConfig.GetCertificate == nil {
		// 서버의 TLS 구성 값에 최소한 하나의 인증서가 포함되어 있지 않거나 GetCertificate 메서드가 nil을 반환하는 경우
		// 매개변수로 입력받은 인증서와 개인키의 경로를 사용하여 파일시스템에서 해당 파일을 읽어서 tls.Certificate 객체를 생성한다.
		cert, err := tls.LoadX509KeyPair(certFn, keyFn)
		if err != nil {
			return fmt.Errorf("loading key pair: %v", err)
		}

		s.tlsConfig.Certificates = []tls.Certificate{cert}
	}

	// 이제 서버에는 클라이언트와의 통신에서 사용할 수 있는 최소한 하나 이상의 인증서를 포함한 TLS 구성이 존재한다.
	// tls.NewListener 함수에 net.Listener 객체와 해당 TLS 구성 정보를 전달하여 TLS를 지원하면 된다.
	// tls.NewListener 함수는 리스너를 받아서 해당 리스너의 Accept 메서드에 TLS를 인지하도록 하는 연결 객체를 반환한다는 점에서 미들웨어처럼 동작한다.
	tlsListener := tls.NewListener(l, s.tlsConfig)

	if s.ready != nil {
		log.Println("before close(s.ready)")
		close(s.ready)
	}

	// ================= 별도의 고루틴에서 리스너로의 연결 요청을 수립하고 처리하여 ServeTLS 메서드 구현을 마무리 한다.
	for {
		conn, err := tlsListener.Accept()
		if err != nil {
			return fmt.Errorf("accept: %v", err)
		}

		go func() {
			for {
				if s.maxIdle > 0 {
					err := conn.SetDeadline(time.Now().Add(s.maxIdle))
					if err != nil {
						return
					}
				}

				buf := make([]byte, 1024)
				n, err := conn.Read(buf)
				log.Println("from client", string(buf[:n]))
				if err != nil {
					return
				}

				_, err = conn.Write(buf[:n])
				log.Println("to client", string(buf[:n]))
				if err != nil {
					return
				}
			}
		}()
	}
}
