/*
대부분의 바이트 슬라이스 []byte 는 문자열을 저장하는 데 사용하기 때문에
Go에서는 바이트 슬라이스를 문자열 타입으로 쉽게 바꿀 수 있도록 지원한다.
다른 타입의 슬라이스와 큰 차이는 없다. 다른 점은 읽거나 쓸 크기를 지정할 수 있으므로
Go가 바이트 슬라이스를 파일 입출력에 사용한다는 점이다.
바이트는 어떤 컴퓨터에서도 동일하게 사용할 수 있는 단위이기 때문이다.

Go는 char 타입이 없으므로 byte나 rune 타입을 사용해야 된다. 하나의 바이트는 하나의 ASCII 문자를 저장할 수 있고
rune은 유니코드 문자를 저장할 수 있다.
*/

package main

import "fmt"

func main() {
	b := make([]byte, 12)
	fmt.Println("Byte slice: ", b)

	b = []byte("Byte slice €") //€ 와 같은 유니코드 문자들은 2바이트 이상의 크기를 차지한다.
	fmt.Println("Byte slice: ", b)
	fmt.Printf("byte slice as a text 1 : %s\n", b)
	fmt.Println("byte slice as a text 2 :", string(b))

	//b의 길이
	fmt.Println("length slice : ", len(b)) // b는 12개의 문자를 가지고 있지만 길이는 14이다.

}
