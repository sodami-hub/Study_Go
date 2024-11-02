// 파일서버 만들기
// 실행파일을 만들어 실행하고
// localhost:3000/gopher.jpg 하면 파일을 불러온다.

package main

import "net/http"

func main() {
	http.Handle("/", http.FileServer(http.Dir("static"))) // 특정 경로에 있는 디렉터리를 파일 서버로 저장
	http.ListenAndServe(":3000", nil)
}
