package main

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	age  int
}

type Students []Student

func (s Students) Len() int {
	return len(s)
}

func (s Students) Less(i, j int) bool {
	return s[i].age < s[j].age
}

func (s Students) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	s := Students{
		{"화랑", 31}, {"백두산", 52}, {"류", 42}, {"켄", 38}, {"송하나", 18},
	}
	sort.Sort(s)
	fmt.Println(s)
}
