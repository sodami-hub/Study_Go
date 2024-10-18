// 자료구조 3 - stack

package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	v *list.List
}

func (s *Stack) Push(value interface{}) {
	s.v.PushBack(value)
}

func (s *Stack) Pop() interface{} {
	value := s.v.Back()
	if value != nil {
		return s.v.Remove(value)
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	value := s.v.Back()
	if value != nil {
		return value
	}
	return nil
}

func NewStack() *Stack {
	return &Stack{list.New()}
}

func main() {
	stack := NewStack()

	for i := 1; i <= 5; i++ {
		stack.Push(i)
	}

	val := stack.Pop()
	for val != nil {
		fmt.Printf("%v ->", val)
		val = stack.Pop()
	}
}
