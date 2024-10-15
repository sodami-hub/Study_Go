// 요소의 삭제

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	idx := 2

	for i := idx; i < len(slice1)-1; i++ {
		slice1[i] = slice1[i+1]
	} // 삭제하고자 하는 요소의 인덱스를 당겨서 지워준다. 이상태로는 [1,2,4,5,6,6] 슬라이스의 크기를 줄여야됨

	//슬라이스의 크기 조정
	slice1 = slice1[:len(slice1)-1]
	fmt.Println(slice1)

	// 위의 코드를 한줄로
	slice2 := []int{1, 2, 3, 4, 5, 6}

	slice2 = append(slice2[:idx], slice2[idx+1:]...)
	// append() 는 첫번째 인자로 값을 추가할 slice자료형을 받고 뒤에는 추가할 값을 나열하는데
	// slice[idx+1:]은 슬라이스이므로 append로 받을 수 없는 자료이다.
	// 이경우 ... 연산자로 slice를 풀어줬다.

	fmt.Println(slice2)
}
