/*
슬라이스에는 내부 배열을 사용해 구현했다.
내부 배열의 길이는 슬라이스의 용량과 같고 포인터가 존재해 슬라이스 원소를 적절한 배열 원소로 연결해준다.

배열과 슬라이스를 연결해서 이 내용을 이해해보자. Go는 슬라이스를 이용해 배열이나 배열의 일부를 참조할 수 있게 해준다.
따라서 슬라이스를 변경하면 참조하는 배열이 영향을 받기도한다. 하지만 슬라이스의 용량이 바뀔 때 배열과의 연결은 끊긴다.
슬라이스의 용량이 바뀌면 내부 배열의 크기도 바뀌어야 하므로 다른 배열에 연결된다.
*/

package main

import (
	"fmt"
)

func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	a := [4]string{"zero", "one", "two", "three"}
	fmt.Println("a: ", a)

	// a[0] 과 S0을 연결
	var S0 = a[0:1]
	fmt.Println(S0)
	// S[0]의 값을 바꿈
	S0[0] = "S0"

	// a[1],a[2] 와 S12 를 연결
	var S12 = a[1:3]
	fmt.Println(S12)
	// a와 연결된 S12의 값을 바꿈
	S12[0] = "S12_0"
	S12[1] = "S12_1"

	// a의 값도 바뀜
	fmt.Println("a : ", a)

	//S0의 용량
	fmt.Println("Capacity of S0 : ", cap(S0), " Length of S0 : ", len(S0))

	//S0에 원소 4개 추가 - S0의 용량이 바뀌므로 내부 배열(a)와의 연결이 끊긴다.
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	fmt.Println(S0)
	a[0] = "-N1"
	fmt.Println(S0)

	// S0의 용량이 바뀐다.
	// a와 같은 내부 배열이 아니다.
	S0 = append(S0, "N4")
	fmt.Println(S0)

	fmt.Println("Capacity of S0 : ", cap(S0), "Length of S0 : ", len(S0))

}
