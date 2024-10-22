// 제네릭과 인터페이스의 차이

package main

import "fmt"

type NodeType1 struct {
	val  interface{}
	next *NodeType1
}

type NodeType2[T any] struct {
	val  T
	next *NodeType2[T]
}

func main() {
	node1 := &NodeType1{val: 1}
	node2 := &NodeType2[int]{val: 2}

	//	var v1 int = node1.val	// 문제 발생  - node1 의 val은 int 타입이 아니라 interface{} 타입이기 때문에 int에 대입할 수 없다.
	var v1 int = node1.val.(int) // interface를 int형으로 캐스팅
	var v2 int = node2.val

	fmt.Println(v1)
	fmt.Println(v2)
}
