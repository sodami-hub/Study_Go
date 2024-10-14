package main

import (
	"fmt"
	"unsafe"
)

type StringHeader struct {
	Data uintptr
	Len  int
}

func main() {
	str1 := "hello"
	str2 := str1

	StringHeader1 := (*StringHeader)(unsafe.Pointer(&str1))
	StringHeader2 := (*StringHeader)(unsafe.Pointer(&str2))

	fmt.Println(StringHeader1)
	fmt.Println(StringHeader2)
}

/*
&{8561194 5}
&{8561194 5}

Data의 주소값이 같다!
*/
