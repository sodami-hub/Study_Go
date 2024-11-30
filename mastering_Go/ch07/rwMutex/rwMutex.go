package main

import (
	"fmt"
	"sync"
	"time"
)

var Password *secret
var wg sync.WaitGroup

/*
반드시 경쟁 상태를 발생할 수 있는 공유 변수의 구조체 내부에서 선언해야 하는 것은 아니다.
그러나 구조체 내부에 선언하는 것이 일반적인 패턴이다. 이는 코드의 가독성을 높이고,
해당 뮤텍스가 보호하는 데이터와 뮤텍스를 함께 관리하기 쉽게 하기 때문이다.
*/
type secret struct {
	RWM      sync.RWMutex
	password string
}

// Change() 함수는 공유 변수 Password를 변경하고 이를 위해 동시에 하나의 쓰기 연산만 허용하는 Lock() 함수를 사용했다.
func Change(pass string) {
	fmt.Println("Change() function")
	Password.RWM.Lock()
	fmt.Println("Change() Locked")
	time.Sleep(3 * time.Second)
	Password.password = pass
	Password.RWM.Unlock()
	fmt.Println("Change() Unlocked")
}

func show() {
	defer wg.Done()
	Password.RWM.RLock()
	fmt.Println("Show function locked!")
	time.Sleep(2 * time.Second)
	fmt.Println("Pass value :", Password.password)
	defer Password.RWM.RUnlock()
}

func main() {
	Password = &secret{password: "admin"}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go show()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("12345")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		Change("sodami")
	}()

	wg.Wait()

	// Direct access to Password.password
	fmt.Println("Current password value:", Password.password)
}
