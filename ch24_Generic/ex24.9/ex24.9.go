// 제네릭을 사용해 만든 유용한 기본 패키지
// slices 패키지의 바이너리 서치
// func BinarySerch[S ~[]E E cmp.Ordered](x S, target E)(int, bool)
// S 는 E의 슬라이스 형태, E는 대소비교가 가능한 모든 타입 - 반환값은 target의 위치(혹은 있어야 되는 위치)와 bool값
package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"alice", "bob", "charlse"} // 정렬된 형태
	n, found := slices.BinarySearch(names, "bob")
	fmt.Println(n, found)
}
