package main

/*
실수 오차
컴퓨터는 2진수 숫자 체계를 사용하기 때문에 실수를 표현할때 2^-1 = 0.5, 2^-2=0.25, 2^-3=0.125 이고,
0.375는 1 * 2^-2 + 1 * 2^-3 으로 표현한다. 하지만 0.375에서 0.001을 더한 값이 0.376을 표현하려고 해도
아무리 작은 2의 마이너스 승수값을 더해도 절대 0.376이란 정확한 값이 나오지 않는다.

*/

// 실수값을 정확히 표현할 수 없기 때문에 아주 작은 오차는 무시하는 방법으로 값을 비교할 수 있다.
// 하지만 최선의 방법은 아니다.  얼마큼의 오차가 무시할만큼 작은 오차인지에대한 정의가 어렵다.
// 가장 간편하고 좋은 방법은 지수부 표현에서 가장 작은 차이인 가장 오른쪽 비트의 1비트 차이만큼 비교하는 것이다. ex4.8
import "fmt"

const epsilon = 0.000001 // 아주 작은 값

func equal(a, b float64) bool {
	if a > b {
		if a-b <= epsilon {
			return true
		} else {
			return false
		}
	} else {
		if b-a <= epsilon {
			return true
		} else {
			return false
		}
	}
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\n", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.0000000000004
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}
