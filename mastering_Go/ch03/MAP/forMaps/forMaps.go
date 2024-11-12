package main

import "fmt"

func main() {
	aMap := make(map[string]string)
	aMap["123"] = "456"
	aMap["key"] = "A value"

	// range는 map에도 사용할 수 있다.
	for k, v := range aMap {
		fmt.Println(k, ":", v)
	}

	v, ok := aMap["123"]
	if ok {
		fmt.Println("있는 키이다.", v)
	} else {
		fmt.Println("없는 키이다.", v)
	}
}
