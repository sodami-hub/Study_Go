package main

import (
	"context"
	"fmt"
)

type aKey string

// ctx.Value()를 이용해 컨텍스트에서 값을 뽑아내 실제로 값이 있는지 확인한다.
func searchKey(ctx context.Context, k aKey) {
	v := ctx.Value(k)
	if v != nil {
		fmt.Println("found value:", v)
		return
	} else {
		fmt.Println("key not found:", k)
	}
}

func main() {
	myKey := aKey("mySecretValue")
	ctx := context.WithValue(context.Background(), myKey, "mySecret")

	searchKey(ctx, myKey)
	searchKey(ctx, aKey("notThere"))

	// context.Background(), context.TODO() 모두 컨텍스트를 생성한다. 그리고 nil이 아닌 빈 Context를 만들지만 두 함수의 목적은 약간 다르다.
	// context.TODO()는 사용하려는 컨텍스트가 확실하지 않을 때만 사용해야 한다.
	// 그러므로 context.TODO()는 컨텍스트를 사용해야 하지만 무엇을 써야 할지 확실하지 않다는 것을 알려준다.
	emptyCtx := context.TODO()
	searchKey(emptyCtx, aKey("notThere"))
}
