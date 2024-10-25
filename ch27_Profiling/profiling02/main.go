package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"
)

var fibMap [65535]int // 피보나치 결과를 저장할 맵

func Fib(n int) int {
	f := fibMap[n]

	if f > 0 {
		return f
	}
	if n == 0 {
		return 0
	} else if n == 1 {
		f = 1
	} else {
		f = Fib(n-1) + Fib(n-2)
	}
	fibMap[n] = f
	return f
}

func main() {
	f, err := os.Create("cpu.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	fmt.Println(Fib(50))

	time.Sleep(5 * time.Second)
}

/*
Fib() 함수가 아예 위쪽에서 사라졌다.
성능이 비약적으로 상승했다.

> go tool pprof .\cpu.prof
File: main.exe
Build ID: C:\Users\leejinhun\goproject\ch27_Profiling\profiling02\main.exe2024-10-25 16:36:31.5799657 +0900 KST
Type: cpu
Time: Oct 25, 2024 at 4:36pm (KST)
Duration: 5s, Total samples = 40ms (  0.8%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 40ms, 100% of 40ms total
Showing top 10 nodes out of 24
      flat  flat%   sum%        cum   cum%
      40ms   100%   100%       40ms   100%  runtime.cgocall
         0     0%   100%       10ms 25.00%  internal/syscall/windows.Module32Next
         0     0%   100%       30ms 75.00%  internal/syscall/windows/registry.Key.GetMUIStringValue
         0     0%   100%       30ms 75.00%  internal/syscall/windows/registry.regLoadMUIString
         0     0%   100%       40ms   100%  runtime/pprof.(*profileBuilder).readMapping
         0     0%   100%       40ms   100%  runtime/pprof.newProfileBuilder
         0     0%   100%       30ms 75.00%  runtime/pprof.peBuildID
         0     0%   100%       40ms   100%  runtime/pprof.profileWriter
         0     0%   100%       30ms 75.00%  sync.(*Once).Do (inline)
         0     0%   100%       30ms 75.00%  sync.(*Once).doSlow
(pprof)

*/
