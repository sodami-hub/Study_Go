package main

import (
	"fmt"
	"hash/fnv"
)

type ComparableHasher interface { // 타입 제한에 메소드까지 추가.
	comparable    // 뭐든 비교(==,!=)가 가능한 타입들을 정의한 인터페이스타입
	Hash() uint32 // 메서드
}

type MyString string // MyString은 비교가 가능한 string 값을 가지고 있으며, Hash() uint32 메서드를 가지고 있어서
//ComparableHasher 타입에 만족한다.

func (s MyString) Hash() uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func Equal[T ComparableHasher](a, b T) bool {
	if a == b {
		return true
	}
	return a.Hash() == b.Hash()
}

func main() {
	var str1 MyString = "Hello"
	var str2 MyString = "World"
	fmt.Println(Equal(str1, str2))
}
