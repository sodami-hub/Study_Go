/*
서버로 데이터를 전송하기 전에 먼저 요청을 수학하고 처리할 핸들러를 작성한다.
아래 코드에서는 User라는 타입을 만들고 이를 이용해 JSON으로 인코딩하여 핸들러에 전송한다.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	First string
	Last  string
}

// handlePostUser() 함수는 POST 요청을 처리.
func handlePostUser(w http.ResponseWriter, r *http.Request) {
	// Go의 HTTP 클라이언트와는 다르게 서버는 반드시 명시적으로 요청 보디를 닫기전에 소비해야만 한다.
	defer func(r io.ReadCloser) {
		_, _ = io.Copy(io.Discard, r)
		_ = r.Close()
	}(r.Body)

	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	var u User
	err := json.NewDecoder(r.Body).Decode(&u) // 요청 보디에 존재하는 JSON을 User 객체로 디코딩한다. 디코딩이 성공하면 상태 코드를 StatusAccepted로 설정한다.
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Decode Failed", http.StatusBadRequest)
		return
	}
	fmt.Println("from client :", u.First, "\t", u.Last)
	w.WriteHeader(http.StatusAccepted)
}

const PORT = ":1234"

func main() {
	mux := http.NewServeMux()

	s := &http.Server{
		Addr:    PORT,
		Handler: mux,
		// IdleTimeout:  10 * time.Second,
		// ReadTimeout:  time.Second,
		// WriteTimeout: time.Second,
	}

	mux.Handle("/postuser", http.HandlerFunc(handlePostUser))

	fmt.Println("server ready! port: 1234, request: /postuser, method: post, body : user data")
	s.ListenAndServe()
}
