/*
타입 단언을 이용해 실제 값을 알아내려고 시도할 때 다음과 같은 두 가지 경우가 있다.

1. 올바른 데이터 타입을 사용하면 아무 이슈 없이 내부의 값을 가져올 수 있다.
2. 올바르지 않은 데이터 타입을 사용하면 프로그램은 패닉이 일어난다.

아래 코드의 마지막부분에서 패닉이 일어나는 이유가 나온다.
map[string]interface{} 에 대해서 mapEmpty.go 에서 살펴보겠다.
*/

package main

import (
	"fmt"
)

// 아래 함수는 int 값을 인터페이스로 감싸서 반환한다.
func returnNumber() interface{} {
	return 12
}

func main() {
	anInt := returnNumber()

	// 인터페이스로 감싸진 변수의 값을 알아냈다.
	number := anInt.(int)
	number++
	fmt.Println(number)

	// 값을 가져올 수 있는 타입 단언을 사용하지 않았으므로 다음 문장은 실패한다.
	// anInt++

	// 다음 문장은 실패하지만, ok bool 변수가 타입 단언이 성공했는지 아닌지 알려준다.
	value, ok := anInt.(int64)
	if ok {
		fmt.Println("Type assertion successful: ", value)
	} else {
		fmt.Println("Type assertion failed!")
	}

	// 다음 문장은 성공적이지만, 타입 단언이 성공했는지 확인하지 않기 때문에 위험하다. 우연히 성공했을 뿐이다.
	i := anInt.(int)
	fmt.Println("i:", i)

	// anInt가 bool이 아니기 때문에 패닉을 일으킨다.
	_ = anInt.(bool) // panic: interface conversion: interface {} is int, not bool
}
