// error은 인터페이스로 문자열을 반환하는 Error() string 메서드로 구성되어 있다.
// 즉 어떤 타입이든 문자열을 반환하는 Error() 메서드를 포함하고 있다면 에러로 사용할 수 있다.
// 이를 이용해서 더 많은 정보를 포함시킬 수 있다.

package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string { // error 인터페이스의 구현체
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error { // error -> 문자열을 반환하는 인터페이스...
	if len(password) < 8 {
		return PasswordError{len(password), 8} // error인터페이스 형으로 return 됨.
		// PasswordError 구조체안에 error인터페이스의 메소드(Error() string)가 구현됨
	}
	return nil
}

func main() {
	err := RegisterAccount("myId", "myPw")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok { // error인터페이스 형으로 반환된 값을 다시 passwordError 형으로 변환
			fmt.Printf("%v Len:%d RequireLen:%d\n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입 완료")
	}
}
