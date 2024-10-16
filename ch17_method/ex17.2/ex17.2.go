// 구조체 뿐아니라 int와 같은 내장 타입들도 별칭 타입을 활용해서 메서드를 가질 수 있다.

package main

import "fmt"

type myInt int

func (a myInt) add(b int) int {
	return int(a) + b // myInt와 int 타입이 다르므로 myInt를 int로 캐스팅
}

func main() {
	var a myInt = 1

	result := a.add(10000)

	fmt.Println(result)

}
