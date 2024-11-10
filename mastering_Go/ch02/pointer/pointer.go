package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory addr of f", &f)

	fp := &f
	fmt.Println("Memory addr of f", fp)
	fmt.Println("Value of f", *fp)

	//f의 값이 바뀐다
	processPointer(fp)
	fmt.Printf("value of f: %.2f\n", f)

	//f의 값은 바꾸지 않는다.
	x := returnPointer(f)
	fmt.Printf("value of f: %.2f\n", f)
	fmt.Printf("value of x: %.2f\n", *x)

	var k *aStructure
	// k가 아무것도 가리키지 않기 때문에 k의 값은 nil이다.
	fmt.Println(k)
	// 따라서 아래와 같은 비교를 할 수 있다.
	if k == nil {
		k = new(aStructure) // k는 nil이 아니고 aStructure의 각 필드는 각 데이터 타입의 제로 값을 가진다.
	}
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil")
	}

}
