// 뮤텍스를 해제하지 않으면 함수를 고루틴 형태로 여러번 실행하면 첫번째 실행을 제외한 나머지 실행에서 뮤텍스가 해제되는 것을 기다리면서 블록된다.
// 예상처럼 아래 코드는 데드락이 발생해 프로그램이 충돌했다. 이런 상황을 피하려면 항상 생성된 뮤텍스를 해제해야 한다.

package main

import (
	"fmt"
	"sync"
)

var m sync.Mutex
var w sync.WaitGroup

func function() {
	m.Lock()
	fmt.Println("Locked!")
}

func main() {
	w.Add(1)
	go func() {
		defer w.Done()
		function()
	}()

	w.Add(1)
	go func() {
		defer w.Done()
		function()
	}()

	w.Wait()
}
