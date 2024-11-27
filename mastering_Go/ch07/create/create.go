package main

import (
	"fmt"
	"time"
)

func printMe(x int) {
	fmt.Println(x)
	return
}

func main() {
	go func(x int) {
		fmt.Printf("%d\n", x)
	}(10)

	go printMe(15)

	time.Sleep(2 * time.Second)
	fmt.Println("exiting...")

}
