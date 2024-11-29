// 정수 값의 제곱을 화면에 출력하는 작업을 한다. 각각의 모든 요청은 하나의 고루틴으로 처리한다.

package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

// 처리할 요청을 관리하는 데 사용할 구조체
type Client struct {
	id      int
	integer int
}

// Client의 결과를 저장한다.
type Result struct {
	job    Client
	square int
}

var size = runtime.GOMAXPROCS(0)
var clients = make(chan Client, size)
var data = make(chan Result, size)
var wg sync.WaitGroup

//clients와 data는 새로운 클라이언트 요청과 결과를 쓸 수 있게 버퍼 채널을 사용했다.
// 프로그램의 성능을 높이고 싶다면 size의 값을 증가시키면 된다.

func worker() {
	for c := range clients {
		square := c.integer * c.integer
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()
}

func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		clients <- c
	}
	close(clients)
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("need #jobs and #workers!")
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	nWorkers, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	go create(nJobs)

	finished := make(chan interface{})

	go func() {
		for d := range data {
			fmt.Printf("Client ID : %d\tint:", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker()
	}

	wg.Wait()
	close(data)

	// finished채널이 닫힐 때까지(출력이 완료될때까지) 블록됨, 서버 과부하 없이 요청을 처리할 수 있다.
	fmt.Printf("Finished: %v\n", <-finished)
}
