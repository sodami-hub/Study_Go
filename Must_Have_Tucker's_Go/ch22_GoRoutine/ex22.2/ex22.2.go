// 서브 고루틴이 종료될 때까지 기다리기

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // waitGroup 객체

func SumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d ~ %d까지의 합은 %d 이다\n", a, b, sum)
	wg.Done() // 작업이 완료됨을 표시
}

func main() {
	wg.Add(10) // 총 작업 개수 설정
	for i := 0; i < 10; i++ {
		go SumAtoB(1, 1000000)
	}
	wg.Wait() // 모든 작업이 끝날때까지 대기
	fmt.Println("모든 작업이 완료됐다.")
}
