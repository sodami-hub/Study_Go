package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type testCases struct {
	method   string
	body     io.Reader
	code     int
	Response string
}

func main() {
	myCase := []testCases{
		{http.MethodGet, nil, http.StatusOK, "Hello, friend!"},
		{http.MethodPost, bytes.NewBuffer([]byte("<world>")), http.StatusOK, "Hello, &lt;world&gt;!"},
		{http.MethodHead, nil, http.StatusMethodNotAllowed, ""},
	}

	client := new(http.Client)
	path := fmt.Sprintf("http://localhost:8081")

	for i, c := range myCase {
		r, err := http.NewRequest(c.method, path, c.body)
		if err != nil {
			fmt.Println("make request error", i, err)
			continue
		}

		resp, err := client.Do(r)
		if err != nil {
			fmt.Println("send req error", err)
			continue
		}
		fmt.Println("expect statusCode :", c.code)
		fmt.Println("actual statusCode :", resp.StatusCode)

		if resp.StatusCode != c.code {
			fmt.Println("status code error")
			continue
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("read body error", err)
			continue
		}
		_ = resp.Body.Close()

		fmt.Println("expect response : ", c.Response)
		fmt.Println("actual respons : ", string(b))

	}

}
