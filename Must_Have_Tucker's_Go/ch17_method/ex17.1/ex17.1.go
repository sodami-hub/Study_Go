package main

import "fmt"

type account struct {
	balance int
}

func withdrawFunc(a *account, amount int) { // 일반 함수 표현
	a.balance -= amount
}

func (a *account) withdrawMethod(amount int) { // 메서드 표현
	a.balance -= amount
}

func main() {
	a := &account{100} // balance가 100인 account 포인터 변수 생성

	// withdrawFunc와 withdrawMethod는 완전히 똑같은 동작한다. 하지만 Func는 함수이고, Method는 Method이다.
	// 따라서 호출 방법이 다르다. 구조체에서 필드가 해당 구조체에 속하듯이 메서드는 해당 리시버 타입에 속한다.
	// 따라서 withdrawMethod()는 리시버 타입인 *account 타입에 속한 메서드이다.

	withdrawFunc(a, 30) // 함수 형태 호출

	a.withdrawMethod(30) // 메서드 형태 호출

	fmt.Printf("%d \n", a.balance)
}
