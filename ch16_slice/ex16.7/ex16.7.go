package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	slice := array[1:2] // slice는 array[1]을 가리키는 len-1, cap-4 의 슬라이스다.
	// 배열을 슬라이싱해서 반환된 슬라이스는 배열의 요소를 카리키는 포인터값을 갖는다.
	// 그래서 서로 영향을 받는다!
	fmt.Println("array : ", array)
	fmt.Println("slice : ", slice, len(slice), cap(slice))
	// 슬라이싱에서 인덱스를 두개만 사용할 때 cap의 크기는 배열(혹은 슬라이스)의 전체 길이에서 시작인덱스를 뺀 값이다.
	// 인덱스를 3개 사용하면 cap까지 조정 가능하다.
	// slice[시작인덱스:끝인덱스:최대인덱스(cap)]

	array[1] = 100

	fmt.Println("after change array[1]")
	fmt.Println("array : ", array)
	fmt.Println("slice : ", slice, len(slice), cap(slice))

	slice = append(slice, 500)

	fmt.Println("after append 500")
	fmt.Println("array : ", array)
	fmt.Println("slice : ", slice, len(slice), cap(slice))

}
