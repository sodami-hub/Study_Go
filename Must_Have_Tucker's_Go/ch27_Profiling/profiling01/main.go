// 특정 구간 프로파일링

package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

func Fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)     // 프로파일 시작!
	defer pprof.StopCPUProfile() // 프로그램 종료 전에 프로파일링 종료

	fmt.Println(Fib(50))

	// 5초를 대기
	time.Sleep(5 * time.Second)
}

/*
Fib() 함수가 전체의 99.75%의 성능을 잡아먹고 있다.
이것만 개선하면 전체 프로그램의 성능이 개선됨

> go tool pprof cpu.prof
File: main.exe
Build ID: C:\Users\leejinhun\goproject\ch27_Profiling\profiling01\main.exe2024-10-25 16:23:46.965946 +0900 KST
Type: cpu
Time: Oct 25, 2024 at 4:25pm (KST)
Duration: 149.23s, Total samples = 140.51s (94.15%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 140.16s, 99.75% of 140.51s total
Dropped 51 nodes (cum <= 0.70s)
      flat  flat%   sum%        cum   cum%
   140.16s 99.75% 99.75%    140.20s 99.78%  main.Fib
         0     0% 99.75%    140.20s 99.78%  main.main
         0     0% 99.75%    140.20s 99.78%  runtime.main
(pprof)

*/
