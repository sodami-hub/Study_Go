// 오차를 없애는 더 나은 방법
// 어떻게 가장 마지막 비트가 1비트만큼 차이나는지 알 수 있을까?
// Go 언어에서는 math 패키지의 Nextafter() 함수를 제공한다.

/*
func Nextafter(x,y float64)(r float64)  // 가장 작은 오차만큼 y를 향해서 더하거나 빼주고, 그 값을 반환한다.
*/

package main

import (
	"fmt"
	"math"
)

func equal(a, b float64) bool {
	return math.Nextafter(a, b) == b
}

func main() {
	var a float64 = 0.1
	var b float64 = 0.2
	var c float64 = 0.3

	fmt.Printf("%0.18f + %0.18f = %0.18f\v", a, b, a+b)
	fmt.Printf("%0.18f == %0.18f : %v\n", c, a+b, equal(a+b, c))

	a = 0.0000000000004 // 아주 작은 값으로 변경
	b = 0.0000000000002
	c = 0.0000000000007

	fmt.Printf("%g == %g : %v\n", c, a+b, equal(a+b, c))
}

// 어디까지나 오차를 무시하는 방법이다! 오차가 작을 뿐이지, 정확한 계산은 아니다.
