// http 쿼리 파라미터 사용하기

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func barHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	name := values.Get("name")
	if name == "" {
		name = "world"
	}
	id, _ := strconv.Atoi(values.Get("id"))
	fmt.Fprintf(w, "hello %s! id:%d", name, id)
}

func main() {
	http.HandleFunc("/bar", barHandler) // 클라이언트 요청에 대한 함수 호출
	http.ListenAndServe(":3000", nil)
	// 두번째 인자에 nil을 넣으면 DefaultServeMux를 사용한다.
	// defaultServerMux 는 HandleFunc() 같은 패키지 함수이다.
}
