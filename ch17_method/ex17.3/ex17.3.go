// 포인터 메서드 vs 값 타입 메서드

package main

import "fmt"

type account struct {
	balance   int
	firstName string
	lastName  string
}

func (a1 *account) withdrawPointer(amount int) {
	a1.balance -= amount
}

func (a2 account) withdrawValue(amount int) {
	//a2.balance -= amount  // call of value 인데 return도 없어서 의미없는 메소드라고 에러표시 떠서 주석 처리
}

func (a3 account) withdrawReturnValue(amount int) account {
	a3.balance -= amount
	return a3 // account 인스턴스를 리턴
}

func main() {
	var mainA *account = &account{100, "joe", "park"}
	mainA.withdrawPointer(30)
	fmt.Println(mainA.balance)

	mainA.withdrawValue(20)
	fmt.Println(mainA.balance)

	var mainB account = mainA.withdrawReturnValue(20)
	fmt.Println(mainB.balance)

	mainB.withdrawPointer(30)
	fmt.Println(mainB.balance)
}
