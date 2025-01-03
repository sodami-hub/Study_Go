package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerWriteHeader(t *testing.T) {
	/*
		아래 핸들러는 400 Bad Request 상태 코드를 생성하고 응답 바디로 bad request를 보내는
		것처럼 보이지만 실제로 그렇게 동작하지 않는다. ResponseWriter의 Writer 메서드를 호출하면
		Go는 암묵적으로 http.StatusOK 상태 코드로 응답의 WriteHeader 메서드를 호출한다.
	*/

	handler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("bad request"))
		w.WriteHeader(http.StatusBadRequest)
	}

	r := httptest.NewRequest(http.MethodGet, "http://test", nil)
	w := httptest.NewRecorder()
	handler(w, r)
	t.Logf("Response status: %q", w.Result().Status)

	handler = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("bad request"))
	}
	r = httptest.NewRequest(http.MethodGet, "http://test", nil)
	w = httptest.NewRecorder()
	handler(w, r)
	t.Logf("Response status: %q", w.Result().Status)
}
