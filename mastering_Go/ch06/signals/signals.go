package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("handleSignal() Caught:", sig)
}

func main() {
	fmt.Printf("Process ID: %d\n", os.Getpid())
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs)

	start := time.Now()
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("Execution time:", duration)
			case syscall.SIGQUIT:
				handleSignal(sig)
				// return 문을 사용하면 고루틴이 종료되므로 return을 사용하면 안된다.
				// 그러나 time.Sleep()은 계속 동작한다.
				os.Exit(0)
			default:
				fmt.Println("Caught:", sig)
			}
		}
	}()

	for {
		fmt.Println("+")
		time.Sleep(10 * time.Second)
	}

}
