package main

import (
	"fmt"
)

func main() {
	var t [5]float64 = [5]float64{23.0, 12.4, 24.2, 24.7, 23.5}

	for i, v := range t { // range 키워드를 사용해서 배열 요소를 순회할 수 있다.
		fmt.Println(i, v)
	}

	for _, v := range t { // 인덱스 값을 사용하고 싶지 않은 경우 _ 을 사용해서 값을 무효화한다.
		fmt.Println(v)
	}
}
