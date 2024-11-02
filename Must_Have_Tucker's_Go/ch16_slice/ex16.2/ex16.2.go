// append()를 상요한 요소 추가.

package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3}

	slice = append(slice, 4)

	fmt.Println(slice)

	slice = append(slice, 500, 600, 700)
	fmt.Println(slice)
}
