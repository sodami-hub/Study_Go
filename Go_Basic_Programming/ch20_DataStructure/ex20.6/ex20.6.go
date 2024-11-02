// map 의 순회

package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

func main() {
	m := make(map[int]Product)

	m[16] = Product{"볼펜", 500}
	m[46] = Product{"지우개", 200}
	m[78] = Product{"자", 1000}
	m[153] = Product{"샤프", 2300}
	m[221] = Product{"샤프심", 400}

	for k, v := range m {
		fmt.Println(k, v)
	}

	//요소의 삭제
	delete(m, 16)
	delete(m, 1) // 없는 요소 삭제 시도 - 아무 동작하지 않음
	fmt.Println(m[46])
	fmt.Println(m[16]) // 요소가 없으면 기본값을 반환 { 0} // string " ", int 0

	// 요소가 있는지 없는지 확인
	v, ok := m[1]
	if !ok {
		fmt.Println("삭제할 값이 없습니다. 없는 키입니다.")
	} else {
		fmt.Println(v)
	}
}
