package main

import (
	"fmt"
	f "fmt"
)

func main() {
	aSlice := []float64{} // 빈 슬라이스를 만든다.

	f.Println(aSlice, len(aSlice), cap(aSlice))

	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -23.134)
	f.Println(aSlice, len(aSlice), cap(aSlice))

	// 용량이 4인 슬라이스
	t := make([]int, 4)
	f.Println(t, len(t), cap(t))
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	//슬라이스에 더이상 남은 공간이 없을 때 append()를 사용해야 된다.
	t = append(t, -5)
	fmt.Println(t)

	//2차원 슬라이스
	// 필요한 만큼의 차원의 가질 수 있다.
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	for _, i := range twoD {
		for _, j := range i {
			f.Print(j, " ")
		}
		f.Println()
	}

	// make를 통해서 2차원 슬라이스를 만들 때는 1차원에 크기를 정의해야 된다.
	make2D := make([][]int, 3)
	f.Println(make2D)
}
