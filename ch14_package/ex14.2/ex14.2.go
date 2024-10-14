package main

import (
	"ch14_package/ex14.2/publicpkg"
	"fmt"
)

func main() {
	fmt.Println("PI : ", publicpkg.PI)
	publicpkg.PublicFunc()

	var myint publicpkg.MyInt = 10
	fmt.Println("myint: ", myint)

	var mystruct = publicpkg.MyStruct{Age: 18}
	fmt.Println("mystruct : ", mystruct)

	fmt.Println(publicpkg.ScreenSize) // 100. 공개되는 변수

}
