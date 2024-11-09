package main

import (
	"fmt"
)

func main() {
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(aSlice)
	l := len(aSlice)

	fmt.Println(aSlice[0:5])
	fmt.Println(aSlice[:5])

	fmt.Println(aSlice[l-2 : l])
	fmt.Println(aSlice[l-2:])

	t := aSlice[0:5:10]
	fmt.Println(t, len(t), cap(t))

	// 슬라이싱의 세번째 요소로 새로 만들어질 슬라이스의 용량을 지정할 수 있다. 하지만 원래 배열 혹은 슬라이스의 용량보다 클 수는 없다.
	// -> 커지게 되면 그 자리에 쓰레기 값이 들어갈 수 있다.
	t = aSlice[2:5:10]
	fmt.Println(t, len(t), cap(t))
}
