package main

import "fmt"

func main() {
	var slice []int

	if len(slice) == 0 { // slice의 길이는 0이다.
		fmt.Println("slice is empty", slice)
	}

	// slice[1] = 10 // 패닉발생(할당되지 않은 메모리 공간에 접근해서 비정상 종료)
	fmt.Println(slice)

	var slice1 = []int{1, 2, 3} // 슬라이스 선언 방법 대괄호 안에 길이가 들어가지 않음
	var slice2 = []int{1, 5: 2, 7: 123}

	fmt.Println(slice1)
	fmt.Println(slice2)

	var array = [4]int{1, 2, 3} // 이건 배열이다.
	fmt.Println("길이가 4인 배열선언 : ", array)

	// make()를 사용한 슬라이스 초기화
	var slice3 = make([]int, 5) // 길이가 5인 정수형 슬라이스 make(자료형, 초기길이(len), 최대길이(capacity))
	fmt.Println(slice3)

	slice3 = append(slice3, 6)
	fmt.Println(slice3)

	fmt.Println("슬라이스 순회 for문")
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("%d ", slice1[i])
	}
	fmt.Println()

	fmt.Println("슬라이스 순회 for문의 range")
	for i, v := range slice1 {
		fmt.Printf("%d : %d - ", i, v)
	}

}
