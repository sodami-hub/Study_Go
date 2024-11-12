package main

import (
	"fmt"
)

func main() {
	aMap := map[string]int{}

	aMap["test"] = 1

	fmt.Println(aMap)
	aMap = nil

	fmt.Println("aMap :", aMap)
	if aMap == nil {
		fmt.Println("nil map!")
		aMap = map[string]int{}
	}
	aMap["test"] = 1
	fmt.Println("2 : ", aMap)
	aMap = nil
	aMap["test"] = 1 // nil인 map에 데이터 넣을 수 없다. 충돌 발생!
	//panic: assignment to entry in nil map
}
