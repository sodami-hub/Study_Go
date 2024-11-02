// 자료구조 1. linked list

package main

import (
	"container/list"
	"fmt"
)

func main() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)
	v.InsertBefore(3, e4) // e4앞에다 노드 삽입
	v.InsertAfter(2, e1)  // e1앞에다 삽입

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value, " ")
	}

	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Println(e.Value, " ")
	}
}
