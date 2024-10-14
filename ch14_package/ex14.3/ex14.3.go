package main

import (
	"ch14_package/ex14.3/exinit" //1. exinit 패키지가 임포트 되면서 초기화가 된다.
	"fmt"
)

func main() { //  exinit 패키지가 초기화 되고 main 함수가 실행된다.
	fmt.Println("main function")
	exinit.PrindD()
}

/*
====
f() d: 4
f() d: 5
init function 6
=== exinit 초기화

main function - 메인함수 실행
d: 6 - exnint.PrintD 실행

*/

/*
package exinit

import "fmt"

var (	// 가장 먼저 전역변수가 초기화 된다. 일반적으로 위에서 아래로 초기화된다.
	a = c + b	// a느 c,b가 초기화 된 후 초기화된다.
	b = f() // b 변수 초기화 d = 4
	c = f() // c 변수 초기화 c = 5 이어서 a가 초기화 되고 a는 9가 된다.
	d = 3 // a,b,c 초기화가 완료되면 d = 5가 된다.
)

func init() { // 전역변수의 초기화가 끝나고 init() 함수가 호출된다. d를 1증가시키고 패키지의 초기화가 끝났다.
	d++
	fmt.Println("init function", d)
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}

func PrindD() {
	fmt.Println("d:", d)
}
*/
