// 뮤텍스를 사용하지 않는 두번 째 방법(첫번째 방법(영역 나누기)은 22장에서)은 채널을 사용해서 역할을 나누는 방법이다.
// 예를들어 자동차 공장에서 자동차를 차체 생산 -> 바퀴 설치 -> 도색 -> 완성 단계라고 가정하자
// 각 공정이 1초 걸린다고 하면 3초가 걸린다.
// 그런데 3명이 하나의 공정을 담당하면 처음에만 3초가 걸리고 그 뒤로는 1초가 걸린다.
// 이런 개념으로 시작

package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Println("start Factory!!")

	wg.Add(3)
	// 3개의 채널을 사용해서 각각의 역할을 수행하도록 한다.
	// 한쪽에서 데이터를 생성해서 넣어주면 다른 쪽에서 생성된 데이터를 빼서 사용하는 방식을 생산자 소비자 패턴이라고 한다!!
	// MakeBody : 생산자 / InstallTire : 소비자 이면서 PaintCar에 대해서는 생산자 / PaintCar : 소비자...
	go MakeBody(tireCh) // 1. 고루틴 생성
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("close the factory")
}

func MakeBody(tireCh chan *Car) { // 차체 생산
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second) // 10초 뒤에 종료
	for {
		select {
		case <-tick:
			// Make a body
			car := &Car{}
			car.Body = "Sport car"
			tireCh <- car
		case <-after: // 10초 뒤에 종료
			close(tireCh)
			wg.Done()
			return
		}
	}
}

func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh {
		time.Sleep(time.Second)
		car.Tire = "HankookTire"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

func PaintCar(paintCh chan *Car) {
	for car := range paintCh {
		time.Sleep(time.Second)
		car.Color = "addalok addalok"
		duration := time.Now().Sub(startTime)
		fmt.Printf("%.2f Complete Car : %s / %s / %s \n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}
