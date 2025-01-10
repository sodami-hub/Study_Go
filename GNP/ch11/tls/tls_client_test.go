package ch11

import (
	"crypto/tls"
	"net"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"golang.org/x/net/http2"
)

func TestClientTLS(t *testing.T) {
	/*
		httptest.NewTLSServer 함수는 HTTPS 서버를 반환한다. 이 코드는 함수 이름 외에도 8장에서 사용한 httptest 패키지와 유사하다.
		아래 코드에서 httptest.NewTLSServer 함수는 새로운 인증서 생성을 포함하여 HTTPS 서버 초기화를 위한 TLS 세부 환경구성까지 처리해 준다.
		신뢰받는 인증 기관으로부터 인증서를 서명하지 않았기에 여느 HTTPS 클라이언트도 생성한 인증서를 신뢰하지 않는다.
	*/
	ts := httptest.NewTLSServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				// 서버에서 HTTP로 클라이언트의 요청을 받으면 요청 객체의 TLS 필드는 nil이 된다. 이러한 케이스를 확인하고, 그에 따라 클라이언트의 요청을 HTTPS로 리다이렉트시킬 수 있다.
				if r.TLS == nil {
					u := "https://" + r.Host + r.RequestURI
					t.Log("u :::::::::::: ", u)
					http.Redirect(w, r, u, http.StatusMovedPermanently)
					return
				}
				t.Log("TLS connecting...")
				w.WriteHeader(http.StatusOK)
			},
		),
	)
	defer ts.Close()

	// 테스트를 위하여 서버 객체의 Client메서드는 서버의 인증서를 신뢰하는 *http.Client 객체를 반환한다. 이 클라이언트를 이용하여 핸들러 내의 TLS와 관련된 코드를 테스트할 수 있다.
	// TLS 연결이 성공하는 케이스
	resp, err := ts.Client().Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status %d; actual status %d",
			http.StatusOK, resp.StatusCode)
	}

	// =========== 새로운 클라이언트로 HTTPS 서버 테스트하기

	/*
		새로운 트랜스포트를 생성하고 TLS 구성을 정의하며 이 트랜스포트를 사용하도록 http2를 구성한 뒤 클라이언트 트랜스포트의 기본 TLS 구성을 오버라이딩한다.
		클라이언트의 TLS 구성 설정의 CurvePreferences 필드 값을 P-256으로 하는 것이 좋다. 이놈은 시간차 공격에 저항이 있으며, TLS 협상에 최소 1.2 이상의 버전을 사용할 것이다.
	*/
	tp := &http.Transport{
		TLSClientConfig: &tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
		},
	}

	// 트랜스포트가 기본 TLS 구성을 사용하지 않기에 클라이언트는 http/2를 기본적으로 지원하지 않는다. http/2를 사용하려면 명시적으로 http/2를 사용하기 위한 함수에 트랜스포트를 전달해 주어야 한다.
	// 이 테스트에서 http/2를 사용하지 않는다. 하지만 트랜스포트의 TLS 구성을 오버라이딩할 경우 http/2 지원이 제거된다는 사실을 잊지 말아야 한다.
	err = http2.ConfigureTransport(tp)
	if err != nil {
		t.Fatal(err)
	}

	client2 := &http.Client{Transport: tp}

	// 이 요청은 클라이언트가 서버가 보내는 인증서의 서명자를 신뢰하지 않기 때문에 에러가 발생한다. -> bad certificate
	_, err = client2.Get(ts.URL)
	if err == nil || !strings.Contains(err.Error(),
		"certificate signed by unknown authority") {
		t.Fatalf("expected unknown authority error; actual: %q ::::::", err)
	}

	// 이를 우회하기 위해서 아래 값을 설정해서 클라이언트의 트랜스포트가 서버의 인증서를 검증하지 않도록 할 수 있다. 디벙깅 외의 목적으로 아래 필드 값을 사용하지 않도록 한다.
	tp.TLSClientConfig.InsecureSkipVerify = true

	resp, err = client2.Get(ts.URL)
	if err != nil {
		t.Fatal("request :::: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expectecd status %d; actual status %d", http.StatusOK, resp.StatusCode)
	}
}

// TCP 상에서의 TLS
func TestClientTLSgoogle(t *testing.T) {
	/*
		tls.DialWithDialer 함수는 *net.Dialer 객체의 포인터와 네트워크 종류(tcp), 네트워트 주소, 그리고 *tls.Config 객체의 포인터를 매개변수로 받는다. 여기서 다이얼러에 30초의 타임아웃과 TLS 설정을 지정해 주었다.
		다이얼이 성공하면 TLS 연결의 세부 상태 정보를 탐색할 수 있다.
	*/
	conn, err := tls.DialWithDialer(
		&net.Dialer{Timeout: 30 * time.Second},
		"tcp",
		"www.google.com:443",
		&tls.Config{
			CurvePreferences: []tls.CurveID{tls.CurveP256},
			MinVersion:       tls.VersionTLS12,
		},
	)
	if err != nil {
		t.Fatal(err)
	}

	state := conn.ConnectionState()
	t.Logf("TLS 1.%d", state.Version-tls.VersionTLS10)
	t.Log(tls.CipherSuiteName(state.CipherSuite))
	t.Log(state.VerifiedChains[0][0].Issuer.Organization[0])

	_ = conn.Close()
}
