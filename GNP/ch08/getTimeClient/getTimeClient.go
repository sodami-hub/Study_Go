package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	resp, err := http.Head("https://vclock.kr/time/")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	now := time.Now().Round(time.Second)
	date := resp.Header.Get("Date")
	if date == "" {
		fmt.Println("no date header received from time.gov")
		return
	}

	dt, err := time.Parse(time.RFC1123, date)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection := resp.Header.Get("Content-Type")
	fmt.Println(connection)

	fmt.Printf("vclock.kr : %s (skew %s)\n", dt, now.Sub(dt))
}
