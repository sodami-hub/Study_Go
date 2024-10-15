// 슬라이스와 배열의 동작 차이
// 단순히 동적.정적 크기 뿐아니라 (쉽게 말하자면)슬라이스는 참조 자료형이다.
/*
type SliceHeader struct {
	Data uintptr  // 실제 배열의 주소값(포인터 변수)
	Len int
	Cap int
}
*/

package main

import "fmt"

func changeArray(arr [5]int) {
	//arr[2] = 200	// 포인터로 변경하라는 표시떠서 주석처리
}

func changeSlice(slice []int) {
	slice[2] = 2000
}

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	var slice = []int{1, 2, 3, 4, 5}

	changeArray(arr)
	changeSlice(slice)

	fmt.Println(arr)
	fmt.Println(slice)
	/*
		[1 2 3 4 5]
		[1 2 2000 4 5]
	*/
}
