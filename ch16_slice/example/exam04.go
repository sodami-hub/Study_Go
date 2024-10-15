package main

import (
	"fmt"
	"sort"
)

type Persons []Person

// Len implements sort.Interface.
func (p Persons) Len() int {
	return len(p)
}

// Less implements sort.Interface.
func (p Persons) Less(i int, j int) bool {
	return p[i].score > p[j].score // 요렇게 하면 score 기준 내림차순 정렬
}

// Swap implements sort.Interface.
func (p Persons) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

type Person struct {
	name        string
	age         int
	score       int
	passSuccess float32
}

func main() {
	s := Persons{
		{"나통키", 13, 45, 78.4},
		{"오맹태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8},
		{"황금산", 16, 36, 89.7},
	}

	sort.Sort(s)
	fmt.Println(s)
}
