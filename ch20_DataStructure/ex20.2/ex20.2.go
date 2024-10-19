// 자료구조 2. 큐 = 링크드 리스트로 구현 뒤로 들어가고 앞으로 빼는 형태로 메소드 구현 FIFO
// Go의 컨테이너 패키지에서 제공됨

package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) {
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} {
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	v := queue.Pop()

	for v != nil {
		fmt.Printf("%v ->", v)
		v = queue.Pop()
	}
}
