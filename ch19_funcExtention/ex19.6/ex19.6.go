package main

import "fmt"

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	i := 0
	for i = 0; i < len(f); i++ {
		f[i] = func() {
			fmt.Println(i) // 리터럴 함수는 i 변수를 캡쳐할 때 i 값이 복사되는 게 아니라 i 변수의 주소를 참조한다.
			//즉 여기 i는 값이 아니라. 외부 변수 i의 주소값이다.
		}
	}

	for i := 0; i < len(f); i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < len(f); i++ {
		v := i // 리터럴 함수에서 출력할 값을 v로 복사. v는 i의 주소값이 아니라 i의 값이다!
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < len(f); i++ {
		f[i]()
	}
}

func main() {
	CaptureLoop()
	CaptureLoop2()
}
