package main

import "fmt"

func main() {
	var a int = 3 // 기본 형태
	var b int     // 초깃값 생략, 초깃값은 타입별 기본값으로 대체(int 초기값 0)
	var c = 4     // 타입 생략, 변수 타입은 우변의 값의 타입이 됨
	d := 5        // 선언 대입문 := 을 사용해서 var 키워드와 타입 생략

	e := 3.14       // e의타입은 float64 타입으로 자동 지정
	s := "hello Go" // s의 타입은 string 으로 자동지정

	fmt.Println(a, b, c, d, e, s)
}

/*
정수형 초깃값 0
실수형 초깃값 0.0
불리언 false
문자열 ""(빈 문자열)
그 외 null(정의되지 않은 메모리 주소를 나타내는 go 키워드)
*/
