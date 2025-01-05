package middleware

import (
	"fmt"
	"net/http"
	"path"
	"strings"
)

func RestrictPrefix(prefix string, next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("request URL : ", r.URL.Path)
			// 요청 URL 경로를 검사
			for _, p := range strings.Split(path.Clean(r.URL.Path), "/") {
				// 주어진 접두사로 시작하는지를 확인
				fmt.Println("p :", p)
				if strings.HasPrefix(p, prefix) {
					// 요청 경로에 주어진 접두사로 시작하는 경우 404 Not Found 상태를 응답한다.
					http.Error(w, "Not Found", http.StatusNotFound)
					return
				}
			}
			next.ServeHTTP(w, r)
		},
	)
}
