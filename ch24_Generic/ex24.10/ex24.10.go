// 제네릭을 사용해서 만든 유용한 페키지
// slices

package main

import (
	"cmp"
	"fmt"
	"slices"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{ // 1. Person은 구조체 타입으로 대소 비교가 되지 않는다. 그래서 Person의 슬라이스 타입또한 BinarySearch를 사용할수 없다.
		{"alice", 55},
		{"bob", 24},
		{"gopher", 13},
	}
	// 대신 BinarySearchFunc 를 사용하면 검색할 수 있다. 세번째 인자로 들어가는 함수를 볼것!
	// BinarySearchFunc[S ~[]E, E, T any](x S, target T, cmp func(E, T) int) (int, bool)
	n, found := slices.BinarySearchFunc(people, "bob", func(a Person, b string) int {
		return cmp.Compare(a.Name, b) // Compare[T Ordered](a, b T) int
	})
	fmt.Println("bob :", n, found)

	// clone

	slice := []int{1, 2, 3, 4, 5}
	slice2 := slices.Clone(slice) // 슬라이스를 복제해서 새로운 슬라이스를 반환.

	fmt.Println(slice2)
	slice2[1] = 100

	fmt.Println(slice, slice2)

	// concat
	fmt.Println(slices.Concat(slice, slice2))
}
