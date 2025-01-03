package handlers

import (
	"fmt"
	"html"
	"html/template"
	"io"
	"net/http"
)

var t = template.Must(template.New("hello").Parse("Hello,{{.}}!"))

func DefaultHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Go의 HTTP 클라이언트에서는 요청 보디를 닫으면 암묵적으로 소비하는 데 반해, 서버에서는 요청 보디는 소비되지 않는다. 확실하게 TCP 세션을 재사용하려면 반드시 요청 보디를 소비해야 한다.
		defer func(r io.ReadCloser) {
			_, _ = io.Copy(io.Discard, r)
			_ = r.Close()
		}(r.Body)

		var b []byte

		switch r.Method {
		case http.MethodGet:
			b = []byte("friend")
		case http.MethodPost:
			var err error
			b, err = io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
		default:
			// "Allow" 헤더가 없기 때문에 RFC 구격을 따르지 않음
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		_ = t.Execute(w, string(b))
	})
}

// methods.go 를 사용한 향상된 핸들러 구현
func DefaultMethodsHandler() http.Handler {
	return Methods{
		http.MethodGet: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { _, _ = w.Write([]byte("Hi, friend!")) }),
		http.MethodPost: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b, err := io.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}
			_, _ = fmt.Fprintf(w, "Hello, %s!", html.EscapeString(string(b)))
		}),
	}
}
