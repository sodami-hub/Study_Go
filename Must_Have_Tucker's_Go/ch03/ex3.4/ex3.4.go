package main

import "fmt"

func main() {
	str := "Hello\tGo\t\tWorld\n\"Go\"is Awesome!\n"

	fmt.Print(str)
	fmt.Printf("%s", str)
	fmt.Printf("%q", str) // 모든 특수 문자가 기능을 잃고 문자 자체로 동작하게 됨.
}
