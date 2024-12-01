package main

import (
	"fmt"
	"time"
)

// 아래 코드는 경쟁 상태가 발생할 수 있다. go루틴의 작동 시점에 따라서 i의 값이 예상한 값과 달라질 수 있다.
// for루프가 끝난 다음에 go루틴이 시작되면 i 값은 21이 된다...
// 이 문제를 해결하기 위해 정확한 값을 캡춰해서 고루틴에 넣어 줄 수 있다.
/*
	for i := 0; i <= 20; i++ {
		go func(i int) {
			fmt.Print(i, " ")
		}(i)
	}

	또는
	for i := 0; i <= 20; i++ {
		i := i
		go func() {
			fmt.Print(i, " ")
		}()
	}

*/
func main() {
	for i := 0; i <= 20; i++ {
		go func(i int) {
			fmt.Print(i, " ")
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println()
}
