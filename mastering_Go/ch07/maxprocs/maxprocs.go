package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("you are using ", runtime.Compiler)
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("using Go version", runtime.Version())
	// rntime.GOMAXPROCS(0)을 호출하면 항상 동시에 실행할 수 있는 CPU최대 개수의 직전 값을 반환한다.
	// 매개변수가 1보다 크거나 가다면 현재 설정을 변경한다. 0은 아무 설정도 변경하지 않는다.
	fmt.Println("GOMAXPROCS :", runtime.GOMAXPROCS(0))
}
