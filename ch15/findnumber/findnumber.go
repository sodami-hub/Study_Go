package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	var stdin = bufio.NewReader(os.Stdin)

	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(100)
	count := 0

	var myNum int

	for {
		fmt.Print("숫자값을 입력하세요 > ")

		_, err := fmt.Scanln(&myNum)

		if err != nil {
			stdin.ReadString('\n')
			continue
		}

		count++

		if num == myNum {
			fmt.Println("정답입니다. 시도한 횟수 : ", count)
			break
		} else if num > myNum {
			fmt.Println("입력하신 숫자가 작습니다.")
		} else {
			fmt.Println("입력하신 숫자가 더 큽니다.")
		}
	}

}
