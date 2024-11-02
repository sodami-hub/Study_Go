// 제네릭의 타입제한과 별칭타입의 선언...

package main

type Integer interface { // 1. 타입 제한
	int8 | int16 | int32 | int64 | int
}

type Integer1 interface { // 1. '~'을 사용하면 해당 타입의 별칭까지 모두 사용 가능하다.
	~int8 | ~int16 | ~int32 | ~int64 | ~int
}

func add[T Integer1](a, b T) T {
	return a + b
}

type Myint int // 별칭타입 정의

func main() {
	add(1, 2) // 정상 실행
	var a Myint = 3
	var b Myint = 4
	add(a, b) // 별칭으로 정의했기때문에 Integer에 포함되지 않아서 에러가 발생 하지만 Integer1으로는 가능
}
