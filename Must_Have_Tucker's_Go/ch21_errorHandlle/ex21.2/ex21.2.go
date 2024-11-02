// fmt 패키지의 Errorf함수를 사용하면 원하는 에러 메시지를 만들 수 있다.
// 또는 errors 패키지의 New() 함수를 이용해서 error를 생성할 수 있다.
package main

import (
	"errors"
	"fmt"
	"math"
)

func Sqrt(f float64) (float64, error, error) {
	if f < 0 {
		return 0, fmt.Errorf("제곱근은 양수여야 한다. f: %g", f), errors.New("error: 에러발생했다")
		//1. f가 음수이면 에러 반환
	}
	return math.Sqrt(f), nil, nil
}

func main() {
	value := -2
	sqrt, err, err2 := Sqrt(float64(value))
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		fmt.Println(err2)
		return
	}
	fmt.Printf("Sqrt(%d) = %v\n", value, sqrt)
}
