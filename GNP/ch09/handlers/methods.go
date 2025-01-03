package handlers

import (
	"io"
	"net/http"
	"sort"
	"strings"
)

type Methods map[string]http.Handler

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
