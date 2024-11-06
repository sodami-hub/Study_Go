// 문자열과 같은 잘못된 입력을 제외하고 입력된 숫자 값들의 최솟값과 최댓값 구하기.

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("인수를 2개 이상 입력하세요.")
	}
	count := 0
	var min, max float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(os.Args[i], 64)

		if err != nil {
			continue
		}

		if count == 0 {
			min = n
			max = n
			count++
			continue
		}

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Max : ", max)
	fmt.Println("Min : ", min)
}
