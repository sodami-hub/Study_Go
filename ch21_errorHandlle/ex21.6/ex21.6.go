// recover() 는 발생한 panic 객체를 반환해준다.

package main

import "fmt"

func f() {
	fmt.Println("f() 함수 시작")
	defer func() { // defer를 사용해 함수 종료 전 함수 리터럴이 실행된다. 전파중인 패닉이 있으면 if문이 실행되고 복구가 된다.
		if r := recover(); r != nil {
			fmt.Println("panic 복구 -", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9/3 =%d\n", h(9, 3))
	fmt.Printf("9/0 =%d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("분모는 0일 수 없다.")
	}
	return a / b
}

func main() {
	f()
	fmt.Println("프로그램이 계속 실행됨") // 패닉이 f() 함수에서 defer를 통해서 복구 됐기 때문에 프로그램이 강제종료되지않고 실행된다. 그래서 마지막 메세지가 출력된다.
}
