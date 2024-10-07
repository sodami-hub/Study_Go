package main

import (
	"fmt"
	"math/big"
)

/*
허용 오차를 줄이는 방법 3. math/big의 Float을 사용하는 방법
*/

func main() {
	a, _ := new(big.Float).SetString("0.1")
	b, _ := new(big.Float).SetString("0.2")
	c, _ := new(big.Float).SetString("0.3")

	d := new(big.Float).Add(a, b)
	fmt.Println(a, b, c, d)
	fmt.Println(c.Cmp(d))
}
