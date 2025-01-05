package main

import (
	"fmt"
	"io"
	"net/http"
)

func drainAndClose(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		fmt.Println("요청 바디를 정리.")
		_, _ = io.Copy(io.Discard, r.Body)
		_ = r.Body.Close()
	})
}

func main() {
	serveMux := http.NewServeMux()

	s := &http.Server{
		Addr:    ":1234",
		Handler: serveMux,
	}

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	serveMux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Hello friend.")
	})
	serveMux.HandleFunc("/hello/there/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "Why, hello there.")
	})

	// 요청을 소비(정리)하는 부분인데...
	drainAndClose(serveMux)

	fmt.Println("Ready to serve at PORT: 1234")
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}

}
