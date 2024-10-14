package main

import "fmt"

func main() {
	var t [5]float64 = [5]float64{23.0, 12.4, 24.2, 24.7, 23.5}

	for i := 0; i < len(t); i++ {
		fmt.Println(t[i])
	}
}
