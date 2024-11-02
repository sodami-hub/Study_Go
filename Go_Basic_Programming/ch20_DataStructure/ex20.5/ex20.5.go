// 자료구조 5 - 맵
// 기본내장객체 : 많이 사용 됨

package main

import (
	"fmt"
)

func main() {
	m := make(map[string]string)
	m["lee"] = "seoul"
	m["kim"] = "suwon"
	m["choi"] = "bucheon"
	m["park"] = "parking"

	m["park"] = "경기도" // 값 변경

	fmt.Printf("lee의 주소는 %s입니다.\n", m["lee"])
	fmt.Printf("park의 주소는 %s입니다.\n", m["park"])

	for k, v := range m {
		fmt.Printf("%s의 주소는 %s 입니다.\n", k, v)
	}
}
