package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var myUrl string
var delay int = 5
var wg sync.WaitGroup

type myData struct {
	r     *http.Response
	error error
}

func connect(c context.Context) error {
	defer wg.Done()
	data := make(chan myData, 1)
	tr := &http.Transport{}
	httpClient := &http.Client{Transport: tr}
	req, _ := http.NewRequest("GET", myUrl, nil)

	go func() {
		response, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			data <- myData{nil, err}
			return
		} else {
			pack := myData{response, err}
			data <- pack
		}
	}()

	select {
	// 컨텍스트가 먼저 타임아웃되면 tr.CancelRequest(req)로 클라이언트 연결이 취소된다.
	case <-c.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("The request was canceled!")
		return c.Err()
	case ok := <-data:
		err := ok.error
		resp := ok.r
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer resp.Body.Close()

		realHTTPData, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Printf("Server Response:%s\n", realHTTPData)
	}
	return nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a URL and a delay!")
		return
	}

	myUrl = os.Args[1]

	if len(os.Args) == 3 {
		t, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
		delay = t
	}

	fmt.Println("Delay :", delay)
	c := context.Background()
	c, cancel := context.WithTimeout(c, time.Duration(delay)*time.Second)
	defer cancel()

	fmt.Println("Connecting to ", myUrl)
	wg.Add(1)
	go connect(c)
	wg.Wait()
	fmt.Println("Exiting...")
}
