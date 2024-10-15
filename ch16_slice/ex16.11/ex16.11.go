// 슬라이스의 중간에 요소 삽입

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}

	idx := 2
	data := 100

	slice = append(slice, 0) // slice 길이를 하나 증가.
	for i := len(slice) - 2; i >= idx; i-- {
		slice[i+1] = slice[i]
	}
	slice[idx] = data

	fmt.Println(slice)

	// 위의 값 삽입 과정을 한줄로.
	slice1 := []int{1, 2, 3, 4, 5, 6}

	slice1 = append(slice1[:idx], append([]int{data}, slice1[idx:]...)...)
	// 안쪽에 만들어진 slice는 불필요한 메모리를 사용한다.
	fmt.Println(slice1)

	// 불필요한 메모리를 사용하지 않도록

	slice2 := []int{1, 2, 3, 4, 5, 6}

	slice2 = append(slice2, 0)
	copy(slice2[idx+1:], slice2[idx:])
	slice2[idx] = data
	fmt.Println(slice2)
}
