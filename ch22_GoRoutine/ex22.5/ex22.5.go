// 데드락!
// 서로가 상대가 가지고 있는 자원을 필요로하는 상황이 발생하면... 데드락 발생... 프로그램 강제 종료

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())

	wg.Add(2)
	fork := &sync.Mutex{} // 포크와 수저 뮤텍스
	spoon := &sync.Mutex{}

	go diningProblem("A", fork, spoon, "포크", "수저") // A는 포크 먼저
	go diningProblem("B", spoon, fork, "수저", "포크")

	wg.Wait()
}

func diningProblem(name string, first, second *sync.Mutex, firstName, secondName string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥을 먹으려 한다.\n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstName)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondName)

		fmt.Printf("%s 밥을 먹는다.\n", name)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
