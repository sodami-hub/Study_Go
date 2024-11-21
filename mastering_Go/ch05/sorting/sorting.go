package main

import (
	"fmt"
	"sort"
)

type Grades struct {
	Name    string
	Surname string
	Grade   int
}

func main() {
	data := []Grades{
		{"J.", "Lewis", 10},
		{"M.", "Tsoukalos", 7},
		{"D.", "Tsoukalos", 8},
		{"L.", "Sodam", 6},
	}

	//SliceIsSorted() 슬라이스가 주어진 함수의 규칙에 따라 정렬됐는지 확인한다.
	isSorted := sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	if isSorted {
		fmt.Println("it is sorted!")
	} else {
		fmt.Println("it is not sorted")
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	fmt.Println(data)
	isSorted = sort.SliceIsSorted(data, func(i, j int) bool {
		return data[i].Grade < data[j].Grade
	})

	if isSorted {
		fmt.Println("it is sorted!")
	} else {
		fmt.Println("it is not sorted")
	}
}
