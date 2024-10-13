/*
레이블을 사용하는 경우 편리할 수 있으나 혼동을 불러일으킬 수 있꼬 자칫 잘못 사용하면 버그가 발생할 수 있다.
되도록 플래그를 사용하도록 하자.
*/

package main

import "fmt"

func main() {
	a := 1
	b := 1

OutterFor: // 레이블 정의
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ {
			if a*b == 45 {
				break OutterFor // 레이블에 가장 먼저 포함된 for 문까지 종료
			}
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
