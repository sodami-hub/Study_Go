package main

import "fmt"

func F(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return F(n-2) + F(n-1)
}

func main() {
	fmt.Println(F(9))
}
