package main

import (
	"fmt"
	"time"
)

func main() {
	str := "2024/11/08 21:31"
	//문자열을 time.Time으로 변환하려면 time.Parse()함수와 함께 포멧문자열을 사용해야 된다.
	//포맷문자열은 다음과 같다.
	// 2006-01-02 15:04:05.999999999 -0700 MST
	//2006은 년, 01은 월, 02는 일, 15는 시간, 04는 분, 05는 초, 999999999는 나노초, -0700은 시간대, MST는 시간대 이름
	parseStr, err := time.Parse("2006/01/02 15:04", str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("type: %T, time: %v\n", parseStr, parseStr)

}
