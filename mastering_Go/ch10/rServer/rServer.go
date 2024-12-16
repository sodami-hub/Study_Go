package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `json:"user"`
	Password string `json:"password"`
}

// 사용자 정보를 갖고 있는 전역변수.. 프로그램 전체에서 공유하므로 동시성 관점에서 안전하지 않다. 하지만 지금은 RESTful 에 대한 이해에 집중한다.
var user User
var DATA = make(map[string]string)

var PORT = ":1234"

// 기본 핸들러의 구현, 실제 서버에서는 기본 핸들러가 서버의 사용법을 출력할 때도 있고, 사용할 수 있는 엔드포인트의 목록을 출력할 수도 있다.
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving : ", r.URL.Path, "/ from ", r.Host)
	w.WriteHeader(http.StatusNotFound)
	Body := "Thanks for visiting!\n"
	fmt.Fprintf(w, "%s", Body)
}

// 현재 날짜와 시간을 반환하는 핸들러. 이런 핸들러는 보통 서버의 상태를 테스트할 때 사용하고 실제 프로덕션 환경에서는 제거된다.
func timeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host)
	t := time.Now().Format(time.RFC1123)
	fmt.Fprintf(w, "The current time is : %s \n", t)
}

// "/add" 엔드포인트
func addHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving :", r.URL.Path, "/ form ", r.Host, r.Method)

	// http.Error() 함수는 에러 메시지 및 HTTP 상태 코드와 함께 클라이언트 요청에 대한 응답을 주는 함수다.
	// 여기서 응답을 보내는 에러 메시지는 일반 텍스트여야 한다.
	// 그러나 클라이언트로 원하는 데이터를 쓰려면 여전히 fmt.Fprintf()를 사용해야 한다.
	if r.Method != http.MethodPost {
		http.Error(w, "Error :", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed!")
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error :", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error :", http.StatusBadRequest)
		return
	}

	if user.Username != "" {
		DATA[user.Username] = user.Password
		log.Println(DATA)
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Error : ", http.StatusBadRequest)
		return
	}
}

// "/get" 엔드포인트
func getHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving : ", r.URL.Path, "/ from :", r.Host, r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Error :", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method Not Allowed!")
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ReadAll - Error :", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal -Error", http.StatusBadRequest)
		return
	}
	fmt.Println(user)

	_, ok := DATA[user.Username]
	if ok && user.Username != "" {
		log.Println("Found!")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s : %s\n", d, DATA[user.Username])
	} else {
		log.Println("NOT FOUND!")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Map - Resource not found!", http.StatusNotFound)
	}
	return
}

// "/delete"
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodDelete {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed")
		return
	}

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ReadAll -Error ", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(d, &user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal - Error ", http.StatusBadRequest)
		return
	}
	log.Println(user)

	_, ok := DATA[user.Username]
	if ok && user.Username != "" {
		if user.Password == DATA[user.Username] {
			delete(DATA, user.Username)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s\n", d)
			log.Println(DATA)
		}
	} else {
		log.Println("User", user.Username, "Not Found!")
		w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Delete - Resource not found!", http.StatusNotFound)
	}
	log.Println("After:", DATA)
	return
}

// "/list"
func listHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving:", r.URL.Path, "from", r.Host, r.Method)
	if r.Method != http.MethodGet {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "%s\n", "Method not allowed")
		return
	}

	if len(DATA) == 0 {
		fmt.Fprintf(w, "%s\n", "have no DATA")
		return
	} else {
		for k, v := range DATA {
			fmt.Fprintf(w, "name: %s || password: %s \n", k, v)
		}
		return
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}

	mux := http.NewServeMux()
	s := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
	}

	mux.Handle("/time", http.HandlerFunc(timeHandler))
	mux.Handle("/add", http.HandlerFunc(addHandler))
	mux.Handle("/get", http.HandlerFunc(getHandler))
	mux.Handle("/delete", http.HandlerFunc(deleteHandler))
	mux.Handle("/", http.HandlerFunc(defaultHandler))
	mux.Handle("/list", http.HandlerFunc(listHandler))

	fmt.Println("Ready to serve at", PORT)
	err := s.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}
