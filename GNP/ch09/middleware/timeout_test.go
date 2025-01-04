package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeoutMiddleware(t *testing.T) {
	/*
		http.TimeoutHandler()는 매개변수로 http.Handler와 타임아웃 시간, 메시지 바디를 매개변수로
		http.Handler를 반환하는 미들웨어이다.

		func TimeoutHandler(h Handler, dt time.Duration, msg string) Handler {
			return &timeoutHandler{
				handler: h,
				body:    msg,
				dt:      dt,
			}
		}

	*/
	// 핸들러가 의도적으로 1분 잠들고 타임아웃을 1초로 설정해서 클라이언트가 응답을 읽는데
	// 시간이 걸리는 것처럼 시뮬레이션해서 htt.Handler가 리턴되지 못하도록 한다.
	handler := http.TimeoutHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			time.Sleep(time.Minute)
		}),
		time.Second,
		"Times out while reading response",
	)

	r := httptest.NewRequest(http.MethodGet, "http://test/", nil)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("unexpected status code : %q", resp.Status)
	}
	t.Log(resp.StatusCode)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()
	actual := string(b)
	if actual != "Times out while reading response" {
		t.Logf("unexpectd body: %q", actual)
	}
	t.Log(actual)
}
