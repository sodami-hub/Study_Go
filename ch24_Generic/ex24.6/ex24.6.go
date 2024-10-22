package main

import "fmt"

// 여러개의 타입이 필요하다면
/*
type Node2[T any, K any] struct {
	val1 T
	val2 K
	next *Node2[T, K]
}
*/

type Node[T any] struct { // 제네릭은 함수뿐 아니라 타입 선언시에도 사용가능 - 구조체에 사용
	val  T
	next *Node[T]
}

func NewNode[T any](v T) *Node[T] {
	return &Node[T]{val: v}
}

func (n *Node[T]) Push(v T) *Node[T] {
	node := NewNode(v)
	n.next = node
	return node
}

func main() {
	node1 := NewNode(1)
	node1.Push(2).Push(3).Push(4)

	for node1 != nil {
		fmt.Print(node1.val, "-")
		node1 = node1.next
	}
	fmt.Println()

	node2 := NewNode("HI")
	node2.Push("hello").Push("bye").Push("goodmorning")

	for node2 != nil {
		fmt.Print(node2.val, "-")
		node2 = node2.next
	}
	fmt.Println()
}
