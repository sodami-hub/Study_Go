package main

import "fmt"

// 타입 변환

func main() {

	/*
		a := 3					// int32
		var b float64 = 3.5		// float64

		var c int =b			// error : float64 변수를 int에 대입 불가
		d := a * b				// error : 다른 타입인 int와 float64의 연산 불가

		var e int64 = 7
		f := a*e				// error : 같은 정수형이지만 int 와 int64 타입이 달라 연산 불가
		var g int = b * 3		// error : 실수가 정수로 자동으로 변환되지 않음
	*/

	a := 3              //int
	var b float64 = 3.5 //float64

	var c int = int(b)  //float64를 int로 타입 변환
	d := float64(a * c) // int를 float64로 타입 변환

	var e int64 = 7
	f := int64(d) * e //  float64를 int64로 바꿔서 연산

	var g int = int(b * 3) // float64 -> int : (3.5 *3) 을 계산하고 int로 변환
	var h int = int(b) * 3 // float64 -> int : 위의 연산과 결과값이 다름 3.5를 3으로 바꾸고 *3을 연산
	fmt.Println(g, h, f)
}
