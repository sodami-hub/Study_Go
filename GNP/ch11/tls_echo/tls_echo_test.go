package tls_echo

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestEchoServerTLS(t *testing.T) {

	// ================ 새로운 TLS 에코 서버 생성 및 백그라운드에서 시작
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	serverAddress := "localhost:34443"
	maxIdle := 5 * time.Second
	server := NewTLSServer(ctx, serverAddress, maxIdle, nil)
	done := make(chan struct{})

	go func() {
		err := server.ListenAndServeTLS("cert.pem", "key.pem")
		if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
			t.Error(err)
			return
		}
		done <- struct{}{}
	}()
	server.Ready()

	log.Println("서버 준비 완료")
	// =================== 서버의 인증서를 클라이언트에 고정하기

	// 먼저 cert.pem 파일을 읽는다.
	cert, err := os.ReadFile("cert.pem")
	if err != nil {
		t.Fatal(err)
	}

	// 새로운 인증서 풀을 생성하고, 인증서를 인증서 풀에 추가한다.
	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(cert); !ok {
		t.Fatal("failed to append certificate to pool")
	}

	// 인증서 풀을 tls.Config 객체의 RootCAs 필드에 추가한다. 이름에서 알 수 있듯이 하나 이상의 신뢰하는 인증서를 인증서 풀에 추가할 수 있다.
	// 이를 이용하면 아직 기존 인증서의 만료 기간이 일부 남은 상황에서 새로운 인증서로 마이그레이션하는 데 유용하게 사용할 수 있다.
	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.CurveP256},
		MinVersion:       tls.VersionTLS12,
		RootCAs:          certPool,
	}
	/*
		위의 구성 정보를 이용하여 생성된 클라이언트는 TLS 협상에서 cert.pem 인증서를 사용한 서버만, 혹은 cert.pem 인증서로 서명된 인증설르 사용한 서버만을 인증한다.
	*/

	// =========================== 고정된 인증서를 사용하여 서버 인증

	conn, err := tls.Dial("tcp", serverAddress, tlsConfig)
	log.Println("클라이언트의 서버 접속")
	if err != nil {
		t.Fatal(err)
	}

	hello := []byte("hello")
	_, err = conn.Write(hello)
	if err != nil {
		t.Fatal(err)
	}

	b := make([]byte, 1024)
	n, err := conn.Read(b)
	if err != nil {
		t.Fatal(err)
	}

	if actual := b[:n]; !bytes.Equal(hello, actual) {
		t.Fatalf("expected %q; actual %q", hello, actual)
	}

	cancel()
	<-done
}
