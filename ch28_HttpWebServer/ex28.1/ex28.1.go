// http 서버 만들기
// "/" 요청에 대한 응답 메소드

package main

import (
	"fmt"
	"net/http"
)

func main() {
	//
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	http.ListenAndServe(":3000", nil)
	// 두번째 인자에 nil을 넣으면 DefaultServeMux를 사용한다.
	// defaultServerMux 는 HandleFunc() 같은 패키지 함수이다.

}
