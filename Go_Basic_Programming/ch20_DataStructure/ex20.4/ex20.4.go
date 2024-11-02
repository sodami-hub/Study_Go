// 자료구조 4 - 링
// Go의 컨테이너 패키지에서 제공됨
// 링의 사용
// 저장할 개수가 고정되고, 오래된 요소는 지워도 되는 경우에 적합
// 1. 실행 취소 기능
// 2. 고정 크기 버퍼 기능
// 3. 리플레이 기능

package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)

	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Next()
	}

	fmt.Println()

	for j := 0; j < n; j++ {
		fmt.Printf("%c ", r.Value)
		r = r.Prev()
	}
}
