/*
중첩 for문과 break, 레이블

중첩 for문에서 break를 사용하면 break가 속한 for문에서만 빠져나온다.
모든 for문을 빠져나가고 싶을 때는 어떻게 해야 할까? 첫 번째 방법은 불리언 변수를 사용하는 것이다.
*/

package main

import "fmt"

func main() {
	a := 1
	b := 1
	found := false
	for ; a <= 9; a++ {
		for b = 1; b <= 9; b++ { // 두번째 for문은 여러번 반복되야 되므로 초깃값을 설정해야 된다.
			fmt.Println(a, " ", b)
			if a*b == 45 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
}
