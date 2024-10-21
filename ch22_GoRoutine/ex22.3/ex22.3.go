// 동시성 문제
// 임계구역에서 경쟁상태 발생!!

package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{0}

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.balance < 0 {
		panic(fmt.Sprintf("balance should not be negative value: %d", account.balance))
	}

	account.balance += 1000
	time.Sleep(time.Millisecond)
	account.balance -= 1000
}
