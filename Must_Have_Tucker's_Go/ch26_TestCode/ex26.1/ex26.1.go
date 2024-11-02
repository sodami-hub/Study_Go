// test 대상 코드

package main

import "fmt"

func square(x int) int {
	fmt.Println(x)
	return 81
}

func main() {
	fmt.Printf("9 * 9 = %d\n", square(9))
}
