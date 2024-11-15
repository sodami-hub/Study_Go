package main

import (
	"fmt"
	"reflect"
)

type Secret struct {
	Username string
	Password string
}

// 다른 구조체를 값으로 포함하는 구조체
type Record struct {
	Field1 string
	Field2 float64
	Field3 Secret
}

func main() {
	A := Record{"String value", -12.123, Secret{"Mihalis", "Tsoukalos"}}

	r := reflect.ValueOf(A) // 구조체의 이름을 가져온다.(포인터 복사?)
	// 그러나 . 연산자로 내부 필드에 접근할 수 없다. Field(인덱스번호?) 로 접근한다.

	fmt.Println("리플렉션해서 받은 r : ", r) // A의 값이 그대로 복사됨

	fmt.Println("원래의 구조체 A : ", A)

	fmt.Println("String value:", r.String())

	iType := r.Type()
	fmt.Printf("i Type: %s\n", iType)
	fmt.Printf("The %d fields of %s are\n", r.NumField(), iType)

	for i := 0; i < r.NumField(); i++ {
		fmt.Printf("\t%s ", iType.Field(i).Name)
		fmt.Printf("\twith type: %s ", r.Field(i).Type())
		fmt.Printf("\tand value _%v_\n", r.Field(i).Interface())
		// Record에 다른 구조체가 임베드돼 있는지 체크한다.
		k := reflect.TypeOf(r.Field(i).Interface()).Kind()
		if k.String() == "struct" {
			fmt.Println(r.Field(i).Type())
		}
		// 위와 같은 내용이지만 내부 표현 형식을 사용한다.
		if k == reflect.Struct {
			fmt.Println(r.Field(i).Type())
		}
	}
}
