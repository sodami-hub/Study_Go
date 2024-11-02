package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()               // 응답에 대한 객체 생성
	req := httptest.NewRequest("GET", "/", nil) // 1. 경로 테스트 // 요청에 대한 객체(요청을 생성)

	mux := MakeWebHandler() //2	// 테스트하려는 핸들러 인스턴스르 가져옴
	mux.ServeHTTP(res, req) // 핸들러 인스턴스에 요청(req), 응답(res)를 넣고 실행

	assert.Equal(http.StatusOK, res.Code)     //3. Code 확인  // 응답의 코드값을 비교
	data, _ := io.ReadAll(res.Body)           //4. 데이터를 읽어서 확인 // 응답의 body의 값을 읽어온다.
	assert.Equal("Hello World", string(data)) // 비교한다.
}

func TestBarHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil) // 5. bar경로 테스트

	mux := MakeWebHandler() //2
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := io.ReadAll(res.Body)
	assert.Equal("Hello Bar", string(data))
}
