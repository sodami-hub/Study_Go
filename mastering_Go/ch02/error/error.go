package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// errors.New() 로 커스텀 메시지 만들기 -> 실제로 error.New()를 사용해서 커스텀할 일은 거의 발생하지 않는다.
func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}
	return nil
}

// fmt.Errorf()로 커스텀 에러 메시지 만들기 -> fmt.Errorf() 를 사용하면 더 자유롭게 출력할 수 있다.
func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID : %d", a, b, os.Getuid())
	}
	return nil
}

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("check() ended normally")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)
	if err.Error() == "this is a custom error message" {
		fmt.Println("custom error detected!")
	}
	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("Int value is", i)
	}

	i, err = strconv.Atoi("y123")
	if err != nil {
		fmt.Println(err)
	}
}
