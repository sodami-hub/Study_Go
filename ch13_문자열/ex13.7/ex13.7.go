package main

import "fmt"

func main() {
	str := "Hello 월드"

	for i := 0; i < len(str); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", str[i], str[i], str[i])
	}
}

/*
string의 인덱스로 접근하면 요소의 타입은 unit8 즉 1바이트이다.

타입:uint8 값:72 문자값:H
타입:uint8 값:101 문자값:e
타입:uint8 값:108 문자값:l
타입:uint8 값:108 문자값:l
타입:uint8 값:111 문자값:o
-> 영문은 바이트 단위로 출력이 됨.

타입:uint8 값:32 문자값:
타입:uint8 값:236 문자값:ì
타입:uint8 값:155 문자값:
타입:uint8 값:148 문자값:
타입:uint8 값:235 문자값:ë
타입:uint8 값:147 문자값:
타입:uint8 값:156 문자값:
-> 한글은 바이트 단위로 출력이 안됨.


어떻게 문자열을 순회할까?
1. []rune 타입사용
2. range 키워드 사용.
*/
