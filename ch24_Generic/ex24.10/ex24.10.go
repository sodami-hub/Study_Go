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
}
