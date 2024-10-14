package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var stdin = bufio.NewReader(os.Stdin)

	money := 1000
	var slotNum int
	var myNum int

	for true {
		fmt.Print("당신의 선택은(1~5) > ")
		_, err := fmt.Scanln(&myNum)
		if err != nil || (myNum < 1 || myNum > 5) {
			fmt.Print("다시입력하세요.")
			stdin.ReadString('\n')
			continue
		}

		slotNum = rand.Intn(5) + 1

		if myNum != slotNum {
			fmt.Println("땡")
			money -= 100
			fmt.Println("현재 잔액은 : ", money)
		} else {
			fmt.Println("딩동댕")
			money += 500
			fmt.Println("현재 잔액은 : ", money)
		}

		if money <= 0 || money >= 5000 {
			fmt.Println("게임종료")
			break
		}
	}
}
