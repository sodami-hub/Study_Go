// 문자열은 불변이다.
// 일부 요소만 변경할 수 없다.

package main

import "fmt"

func main() {
	var str string = "hello 월드"
	var slice []byte = []byte(str) // 슬라이스 타입으로 변환

	slice[2] = 'a'

	fmt.Println(str)          // hello world
	fmt.Printf("%s\n", slice) //healo world

}

/*
str과 slice가 가리키는 메모리 공간이 다르다.
[]byte 슬라이스 타입으로의 변환은 복사해서 새로운 메모리 공간을 만든다. 문자열의 불변 성질을 지키기 위함이다.

문자열 연산이 빈번한 경우 메모리 낭비가 생기기 때문에 strings 패키지의 Builder를 이용해서 메모리 낭비를 줄일 수 있다.
*/
