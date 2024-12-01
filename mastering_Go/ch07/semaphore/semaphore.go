package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.org/x/sync/semaphore"
)

var Workers = 4
var sem = semaphore.NewWeighted(int64(Workers))

func worker(n int) int {
	square := n * n
	time.Sleep(time.Second)
	return square
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need num of jobs!")
		return
	}

	nJobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// 여기에 결과를 저장한다.
	var results = make([]int, nJobs)

	//Acquire()에서 필요하다.
	ctx := context.TODO()
	for i := range results {
		// 가중치 1을 사용해서 코드를 실행 - 가중치 1짜리 4개(고루틴 4개) 사용 가능
		err = sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Println("cannot acquire semaphore:", err)
			break
		}

		go func(i int) {
			defer sem.Release(1) // 끝나면 가중치 반환
			temp := worker(i)
			results[i] = temp
		}(i)
	}

	// 모든 토큰을 획득해 sem.Acquire() 호출이 모든 워커/고루틴들이 끝날 때까지 기다린다. 이는 Wait() 호출과 유사하다.
	// 세모포어의 가중치(여기서는 최대가중치)가 모일 때까지 블록!
	err = sem.Acquire(ctx, int64(Workers))
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range results {
		fmt.Println(k, "->", v)
	}
}
