package main

import "fmt"

func main() {
	str := "hello 월드"
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		fmt.Printf("타입:%T 값:%d 문자값:%c\n", runes[i], runes[i], runes[i])
	}

}

/*
타입:int32 값:104 문자값:h
타입:int32 값:101 문자값:e
타입:int32 값:108 문자값:l
타입:int32 값:108 문자값:l
타입:int32 값:111 문자값:o
타입:int32 값:32 문자값:
타입:int32 값:50900 문자값:월
타입:int32 값:46300 문자값:드
*/
