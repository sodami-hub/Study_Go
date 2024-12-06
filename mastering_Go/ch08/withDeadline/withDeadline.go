package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"
)

var timeout = time.Duration(time.Second)

func Timeout(network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a URL")
		return
	}

	if len(os.Args) == 3 {
		temp, err := strconv.Atoi(os.Args[2])
		if err == nil {
			timeout = time.Duration(time.Duration(temp) * time.Second)
		}
	}
	fmt.Println("Timeout value:", timeout)

	URL := os.Args[1]

	// http.Transport의 Dial 필드에 Timeout() 함수를 적용한다.
	t := http.Transport{
		Dial: Timeout,
	}

	// http.Client는 위에서 정의한 http.Transport를 사용하고,
	// http.Transport에서 정의한 Timeout이 사용된다.
	client := http.Client{
		Transport: &t,
	}

	data, err := client.Get(URL)
	if err != nil {
		fmt.Println("Get : ", err)
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
		}
	}
}
