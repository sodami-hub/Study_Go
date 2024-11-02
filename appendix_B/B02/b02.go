package main

import (
	"fmt"
	"goproject/appendix_B/B02/bankaccount"
)

func main() {
	account := bankaccount.NewAccount()
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
