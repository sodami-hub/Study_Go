package main

import "fmt"

func main() {
	for i := 5; i >= 1; i-- {
		for j := i; j >= 1; j-- {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
