// os.Args에서 유효한 입력과 유효하지 않은 입력 구분하기

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	var total, nInts, nFloats int
	errors := make([]string, 0)

	for _, k := range args[1:] {
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		errors = append(errors, k)
	}

	if len(errors) > total {
		fmt.Printf("#read : %d  #ints : %d #floats : %d\n", total, nInts, nFloats)
		fmt.Println("Too much invalid input : ", len(errors))
		for _, k := range errors {
			fmt.Print(k, " ")
		}
	} else {
		fmt.Printf("#read : %d  #ints : %d #floats : %d\n", total, nInts, nFloats)
	}
}
