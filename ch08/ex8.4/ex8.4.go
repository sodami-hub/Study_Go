// switch - case : 비굣값에 true를 넣고 case의 조건문이 true가 될 때 실행한다.
package main

import "fmt"

func main() {
	temp := 18

	switch true {
	case temp < 10, temp > 30:
		fmt.Println("바깥 활동 하지 말자")
	case temp >= 10 && temp < 20:
		fmt.Println("약간 추울 수 있다.")
	case temp >= 15 && temp < 25:
		fmt.Println("좋다")
	default:
		fmt.Println("조으다.")
	}
}
