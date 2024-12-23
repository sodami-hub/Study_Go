/*
[/unlimitBlock.go] 를 실행해서 서버를 띄우고
아래 코드를 실행하면 5초뒤에 타임아웃이 발생한다.
*/

package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var PORT = ":1234"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:1234/block", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if !errors.Is(err, context.DeadlineExceeded) {
			fmt.Println(err)
			return
		}
		fmt.Println("timeout!!")
		return
	}

	defer resp.Body.Close()

}
