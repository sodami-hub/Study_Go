// 같은 자원을 여러 고루틴이 접근하기 때문에 발생하는 문제(경쟁상태)를 해결하기 위한 방법
// 1. 영역을 나누는 방법
// 2. 역할을 나누는 방법

// 영역을 나누는 방법
package main

import (
	"fmt"
	"sync"
	"time"
)

type Job interface { // Job 인터페이스
	Do()
}

type SquareJob struct {
	index int
}

func (j *SquareJob) Do() {
	fmt.Printf("%d 작업 시작\n", j.index) // 각각의 작업
	time.Sleep(1 * time.Second)
	fmt.Printf("%d 작업 완료 - 결과: %d\n", j.index, j.index*j.index)
}

func main() {
	var jobList [10]Job

	for i := 0; i < 10; i++ {
		jobList[i] = &SquareJob{i}
	}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ { // 10가지 작업 할당 - 각 고루틴은 할당된 작업만 한다. 간섭이 발생하지 않는다. 그래서 뮤텍스가 필요 없다.
		job := jobList[i]
		go func() {
			job.Do()
			wg.Done()
		}()
	}
	wg.Wait()
}
