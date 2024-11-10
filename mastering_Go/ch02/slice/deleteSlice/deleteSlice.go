package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("need an integer value.")
		return
	}

	index := arguments[1]

	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("using index :", i)

	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("original slice", aSlice)

	// i번째 인덱스 삭제

	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("after delete : ", aSlice)
}
