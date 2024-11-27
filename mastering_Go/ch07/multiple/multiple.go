package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args
	capa := 10
	if len(args) == 2 {
		cap, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println(err)
			return
		}
		capa = cap
	}

	fmt.Printf("Going to create %d goroutines\n", capa)
	for i := 0; i < capa; i++ {
		go func(x int) {
			fmt.Printf("%d  ", x)
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("\n\n exit...")

}
