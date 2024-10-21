package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	ctx := context.WithValue(context.Background(), "number", 9)
	// 컨텍스트에 값을 추가
	go square(ctx)

	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("number"); v != nil { // Value(a interface{}) interface{} -> 형변환 필요
		n := v.(int)
		fmt.Printf("square : %d\n", n*n)
	}
	wg.Done()
}
