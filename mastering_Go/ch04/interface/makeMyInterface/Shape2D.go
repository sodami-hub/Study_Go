/*
나만의 인터페이스 만들기
1. 새 인터페이스 만들기
2. 기존 인터페이스 합성하기
3. 3차원 형태에서 sort.Interface 구현하기

*/

package main

import (
	"fmt"
	"math"
)

type Shape2D interface {
	Perimeter() float64
}

type circle struct {
	R float64
}

func (c circle) Perimeter() float64 {
	return 2 * math.Pi * c.R
}

func main() {
	cir := circle{2.5}
	s2D := Shape2D(cir)

	fmt.Printf("Perimeter : %.2f\n", s2D.Perimeter())

	var b interface{} = circle{4}
	fmt.Printf("b : %v,  type : %T\n", b, b)
	bb, ok := b.(circle)
	if ok {
		fmt.Println(bb.Perimeter(), bb.R)
	}

	typeCh, ok := b.(Shape2D)
	if ok {
		fmt.Printf("둘레 : %.2f\n", typeCh.Perimeter())
	}

	a := circle{R: 1.5}
	fmt.Printf("R %.2f -> Perimeter %.3f \n", a.R, a.Perimeter())

	_, ok = interface{}(a).(Shape2D)
	if ok {
		fmt.Println("a is a Shape2D!")
	}

}
