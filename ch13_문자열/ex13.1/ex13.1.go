package main

import "fmt"

func main() {
	str1 := "Hello\t'World'\n"

	str2 := `Go is "awesome"!\nGo is simple and\t'powerful'`
	str3 := "Go is \"awesome\"!\nGo is simple and\t'powerful'"

	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(str3)

	/*
		Hello   'World'

		Go is "awesome"!\nGo is simple and\t'powerful'

		Go is "awesome"!
		Go is simple and        'powerful'
	*/

	//백쿼트로 여러줄 문자열 출력하기
	str4 := `죽는 날까지 하늘을 우러러
한 점 부끄럼이 없기를,
잎새에 이는 바람에도
나는 괴로워했다.`
	/*
	   죽는 날까지 하늘을 우러러
	   한 점 부끄럼이 없기를,
	   잎새에 이는 바람에도
	   나는 괴로워했다.
	*/
	fmt.Println(str4)
}
