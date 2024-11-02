package bankaccount

type Account interface {
	Withdraw(money int) int
	Deposit(money int)
	Balance() int
}

func NewAccount() Account { // 계좌 생성함수 -> 인터페이스를 반환한다.
	return &innerAccount{balance: 1000} // 실제 Account 인터페이스의 구현체를 반환하지만 외부에 공개되는 구조체가 아니기 때문에 필드에 접근할 수 없고 인터페이스 메서드로만 사용할 수 있다.
}

type innerAccount struct { // 실제 예금 계좌는 공개되지 않는다.
	balance int
}

func (in *innerAccount) Withdraw(money int) int {
	in.balance -= money
	return in.balance
}

func (in *innerAccount) Deposit(money int) {
	in.balance += money
}

func (in *innerAccount) Balance() int {
	return in.balance
}
