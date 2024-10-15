package main

import "fmt"

// 슬라이스를 복제하는 방법

func main() {
	// 방법 1
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice2[2] = 10000
	fmt.Println("slice2[2]=10000")
	fmt.Println(slice1)
	fmt.Println(slice2)

	// 방법 2
	slice3 := []int{10, 11, 12, 13, 14}
	slice4 := append([]int{}, slice3[0], slice3[1], slice3[2], slice3[3], slice3[4])

	fmt.Println(slice3)
	fmt.Println(slice4)

	// 방법 3은 다음 예제로
}
