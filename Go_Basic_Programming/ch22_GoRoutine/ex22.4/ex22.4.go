package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex // 패키지 전역 변수 뮤텍스

type Account struct {
	Balance int
}

func DepositAndWithdraw(account *Account) {
	mutex.Lock()         // 뮤텍스 획득
	defer mutex.Unlock() // defer 를 사용한 뮤텍스 반환  - 한 번 획득한 뮤텍스는 반드시 Unlock()을 호출해서 반납해야 된다.
	if account.Balance < 0 {
		panic(fmt.Sprintf("balance should not be negative value: %d", account.Balance))
	}

	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	wg.Add(10)

	accounts := &Account{0}

	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(accounts)
			}
			wg.Done() // 절대 끝나지 않음 - 잔고가 0아래로 떨어질 수 없다.
		}()
	}
	wg.Wait()
}
