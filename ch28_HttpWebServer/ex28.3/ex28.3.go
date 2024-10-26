// ServerMux 인스턴스 이용하기

package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	//HandleFunc은  이름은 같지만 Http.HandleFunc()함수는 DefaultServeMux 핸들러에 등록하지만
	//mux.HandleFunc()메서드(ServeMux 인스턴스의 메서드)는 ServeMux인스턴스에 핸들러를 등록한다.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello world")
	})
	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello banana")
	})
	http.ListenAndServe(":3000", mux) // ServeMux를 사용해서 다양한 요청에 대한 응답을 작성할 수 있다.
}
