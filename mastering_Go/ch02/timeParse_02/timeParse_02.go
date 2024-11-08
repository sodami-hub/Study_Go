package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	if len(os.Args) != 2 {
		fmt.Println("Usage: dates parse_string")
		return
	}

	dateString := os.Args[1]

	// 날짜만 존재하는가?
	d, err := time.Parse("02 January 2006", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	}

	// 날짜 + 시간 값인가??
	d, err = time.Parse("02 January 2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	//숫자 형식으로 날짜와 시간이 표현돼 있는가?
	d, err = time.Parse("02-01-2006 15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Date:", d.Day(), d.Month(), d.Year())
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	// 시간만 있는 포맷인가?
	d, err = time.Parse("15:04", dateString)
	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute())
	}

	//유닉스 에포크 시간 다루기
	t := time.Now().Unix()
	fmt.Println("Epoch time:", t)
	// Epoch 시간을 time.Time 값으로 변환하기
	d = time.Unix(t, 0)
	fmt.Println("Date:", d.Day(), d.Month(), d.Year())
	fmt.Printf("Time: %d:%d\n", d.Hour(), d.Minute())
	duration := time.Since(start)
	fmt.Println("execution time:", duration)
}

/*
$ go run timeParse_02.go "08-11-2024 01:10"
Full: 2024-11-08 01:10:00 +0000 UTC
Date: 8 November 2024
Time: 1 10
Epoch time: 1731078625
Date: 9 November 2024
Time: 0:10
execution time: 143.715µs
*/
