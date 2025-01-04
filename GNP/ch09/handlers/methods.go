package handlers

import (
	"io"
	"net/http"
	"sort"
	"strings"
)

type Methods map[string]http.Handler

/*
ServeHTTP() 메서드는 Go언어의 http.Handler 인터페이스의 메서드로, 클라이언트의 HTTP 요청이 들어오면 자동으로 호출된다.

	type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
	}

따라서 Methods 타입에는 ServeHTTP(http.ResponseWriter, *http.Request) 메서드가 구현됐기 때문에
타입 Methods는 http.Handler 인터페이스의 구현체이다.
*/
func (h Methods) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(r io.ReadCloser) {
		_, _ = io.Copy(io.Discard, r)
		_ = r.Close()
	}(r.Body)

	if handler, ok := h[r.Method]; ok {
		if handler == nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		} else {
			handler.ServeHTTP(w, r)
		}
		return
	}

	w.Header().Add("Allow", h.allowedMethods()) // 사용가능한 HTTP메서드를 나열해서 응답해줌
	if r.Method != http.MethodOptions {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h Methods) allowedMethods() string {
	a := make([]string, 0, len(h)) // make(타입, 기본길이, 기본 최대 용량)

	for k := range h {
		a = append(a, k)
	}
	sort.Strings(a)

	return strings.Join(a, ",")
}
