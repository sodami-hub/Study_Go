package main

import (
	"fmt"
	"runtime/metrics"
	"sync"
	"time"
)

func main() {
	// nGo 변수는 수집하고자 하는 메트릭의 경로를 저장한다.
	const nGo = "/sched/goroutines:goroutines"

	// 메트릭의 샘플을 가져오기 위한 슬라이스 metrics.Sample의 필드는 Name, Value 이다.
	getMetric := make([]metrics.Sample, 1)
	getMetric[0].Name = nGo

	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(4 * time.Second)
		}()
		// 실제 데이터를 가져온다.
		metrics.Read(getMetric) // metrics.Read() 함수는 getMetric 슬라이스로 주어진 메트릭스들을 수집한다.
		if getMetric[0].Value.Kind() == metrics.KindBad {
			fmt.Printf("metric %q no longer supported\n", nGo)
		}

		// 원하는 메트릭을 읽은 다음 프로그램에서 사용할  수 있게 숫자 값으로 변경한다.
		mVal := getMetric[0].Value.Uint64()
		fmt.Printf("Number of goroutines: %d\n", mVal)

	}

	wg.Wait()

	metrics.Read(getMetric)
	mVal := getMetric[0].Value.Uint64()
	fmt.Printf("Before exiting: %d\n", mVal)
}

/*
Number of goroutines: 2
Number of goroutines: 3
Number of goroutines: 4
Before exiting: 1

전체 고루틴의 수는 4개이고 프로그램이 끝나기 전에는 main 고루틴 한개만 존재한다.
*/
