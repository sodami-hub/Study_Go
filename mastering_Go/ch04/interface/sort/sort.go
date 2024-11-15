package main

import (
	"fmt"
	"sort"
)

type S1 struct {
	F1 int
	F2 string
	F3 int
}

// S2구조체는 F3.F1값을 기반으로 정렬하고 싶다.
type S2 struct {
	F1 int
	F2 string
	F3 S1
}

type S2slice []S2

// ====== sort 인터페이스의 구현
func (s2 S2slice) Len() int {
	return len(s2)
}
func (s2 S2slice) Less(i, j int) bool {
	if s2[i].F3.F1 < s2[j].F3.F1 {
		return true
	} else {
		return false
	}
}
func (s2 S2slice) Swap(i, j int) {
	s2[i], s2[j] = s2[j], s2[i]
}

// ============================

func main() {
	data := []S2{
		{1, "One", S1{1, "S1_1", 10}},
		{2, "Two", S1{2, "S1_1", 20}},
		{-1, "Two", S1{-1, "S1_1", -20}},
	}
	fmt.Println("Before:", data)
	sort.Sort(S2slice(data)) // sort.Interface를 구현했기 때문에 Sort를 호출할 수 있다.
	fmt.Println("After:", data)

	// Reverse sorting works automatically
	sort.Sort(sort.Reverse(S2slice(data)))
	fmt.Println("Reverse:", data)
}
