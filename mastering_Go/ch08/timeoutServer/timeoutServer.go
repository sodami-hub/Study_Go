package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served: %s\n", r.Host)
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format(time.RFC1123)
	Body := "The current time is:"
	fmt.Fprintf(w, "<h1 align=\"center\">%s</h1>", Body)
	fmt.Fprintf(w, "<h2 align=\"center\">%s</h2>\n", t)
	fmt.Fprintf(w, "Serving: %s\n", r.URL.Path)
	fmt.Printf("Served time for: %s\n", r.Host)
}

func main() {
	PORT := ":8001"
	arguments := os.Args
	if len(arguments) != 1 {
		PORT = ":" + arguments[1]
	}
	fmt.Println("Using port number: ", PORT)

	m := http.NewServeMux()

	/*
		타임아웃 주기를 정의한다. 타임아웃 주기는 읽기와 쓰기 모두에 정의할 수 있다. ReadTimeout 필드의 값은 전체 클라이언트 요청에서 최대 기간을 의미하고
		이는 본문을 포함한 요청 전체를 읽는 데 걸리는 최대 시간을 의미한다.
		WriteTimeout 필드는 클라이언트가 응답을 보내는 데 허용한 최대 기간을 의미한다.
	*/
	srv := &http.Server{
		Addr:         PORT,
		Handler:      m,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	m.HandleFunc("/time", timeHandler)
	m.HandleFunc("/", myHandler)

	err := srv.ListenAndServe()
	if err != nil {
		fmt.Println(err)
		return
	}
}

/*
위 코드는 별다른 동작을 하지 않는다. 하지만 클라이언트가 아무 요청없이 연결을 맺는다면 클라이언트의 연결은 3초 뒤에 끊어질 것이다.
클라이언트가 서버 응답을 받는 데 3초 이상이 걸릴 때도 마찬가지다.
*/
