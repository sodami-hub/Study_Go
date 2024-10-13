package main

import "fmt"

func main__() {
	for i := 1; i <= 9; i++ {

		for j := 2; j <= 9; j++ {
			if j >= 3 && j <= 6 {
				continue
			}
			fmt.Printf("%-2d * %2d = %3d  ", j, i, i*j)
		}
		fmt.Println()
	}
}
