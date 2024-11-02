package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Student struct {
	Name string
	Age int
	Score int
}

func MakeWebHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/student",StudentHandler)
	return mux
}

func StudentHandler(w http.ResponseWriter, r *http.Request) {
	var student = Student{"aaa",15,76}
	data, _ := json.Marshal(student) // Student 객체를 []byte로 변환
	w.Header().Add("content-type","application/json") // response header에 data(content) type이 json임을 전달
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w,string(data)) // 결과 전송
}

func main(){
	http.ListenAndServe(":3000",MakeWebHandler())
}
