package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

// read()함수는 readValue 채널에서 값을 읽으려고 시도함으로써
// monitor() 함수에 값을 넣어 달라는 시그널을 보내는 것으로 이해할 수 있다.
func read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("give an integer!")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var wg sync.WaitGroup

	for r := 0; r < n; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(10 * n))
		}()
	}

	wg.Wait()
	fmt.Printf("\nlast value: %d\n", read())

}
