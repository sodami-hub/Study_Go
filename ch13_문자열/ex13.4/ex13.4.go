package main

import "fmt"

func main() {
	str1 := "가나다라마"
	str2 := "abcde"

	//len() 내장 함수는 문자열이 차지하는 메모리 크기를 알 수 있다.
	fmt.Printf("%d\n", len(str1)) // 15 - 한글 한글자 3바이트
	fmt.Printf("%d\n", len(str2)) // 5  - 영문 한글자 1바이트

	//글자수를 알아내는 방법
	// rune 슬라이스 타입인 []rune 타입을 사용!
	// string 타입은 연속된 바이트 메모리라면 []rune 타입은 글자들의 배열로 이루어져 있다.
	// 완전히 다른 타입이지만 편의를 위해서 Go 언어는 둘의 상호 타입변환을 지원한다.
	str3 := "ab가나다cde라마"
	runes := []rune(str3)
	fmt.Printf("%d", len(runes)) // 10
}
