// go 에서 c언어 호출

package main

import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	Seed(1)
	fmt.Println(Random())
}
