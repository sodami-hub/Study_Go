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
	// 대신 BinarySearchFunc 를 사용하면 검색할 수 있다.
	n, found := slices.BinarySearchFunc(people, "bob", func(a Person, b string) int {
		return cmp.Compare(a.Name, b)
	})
	fmt.Println("bob :", n, found)
}
