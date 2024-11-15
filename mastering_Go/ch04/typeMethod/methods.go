/*
두 행렬을 더하는 알고리듬을 수행하는 함수와 메서드, 돌일한 코드이지만 함수의 경우 배열을 반환하고 메서드는 결과를 변수에 저장하는 차이가 있다.
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

type ar2x2 [2][2]int

// 일반적인 Add() 함수
func Add(a, b ar2x2) ar2x2 {
	c := ar2x2{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = a[i][j] + b[i][j]
		}
	}
	return c
}

// 타입 메서드 Add()
func (v *ar2x2) Add(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			v[i][j] = v[i][j] + b[i][j]
		}
	}
}

// 타입 메서드 빼기
func (v *ar2x2) Subtract(b ar2x2) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			v[i][j] = v[i][j] - b[i][j]
		}
	}
}

func main() {
	if len(os.Args) != 9 {
		fmt.Println("need 8 integer")
		return
	}

	k := [8]int{}
	for index, i := range os.Args[1:] {
		v, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println(err)
			return
		}
		k[index] = v
	}
	a := ar2x2{{k[0], k[1]}, {k[2], k[3]}}
	b := ar2x2{{k[4], k[5]}, {k[6], k[7]}}

	fmt.Println("Traditional a+b", Add(a, b))
	a.Add(b)
	fmt.Println("a+b", a)
	fmt.Println("a :", a)
	a.Subtract(b)
	fmt.Println("a-b", a)
	fmt.Println("a :", a)
}
