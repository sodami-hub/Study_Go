package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	buffer := []byte("data to write\n")

	// ------------------------------------------------------------------- fmt 페키지의 Fprintf 함수 사용
	// 파일이 있다면 os.Create를 호출하면 해당 파일의 내용을 전부 삭제하고 새로 만든다.
	f1, err := os.Create("/tmp/f1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f1.Close()

	//문자열 변수를 받아서 원하는 파일에 데이터를 기록할 수 있게 도와주고, 기록하고자 하는 io.Writer를 매개변수로 넘기면 된다.
	// 이경우 *os.File 변수가 io.Writer 인터페이스를 구현하므로 작업을 수행할 수 있다.
	fmt.Fprintf(f1, string(buffer))

	// -------------------------------------------------------------------  *os.File 의 WriteString 메서드 사용
	f2, err := os.Create("/tmp/f2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(buffer))
	fmt.Printf("wrote %d bytes\n", n)

	// ------------------------------------------------------------------- bufio 패키지의 NewWriter함수 사용
	f3, err := os.Create("/tmp/f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f3.Close()

	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(buffer))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	//------------------------------------------------------------------- io 패키지의 WriteString 함수 사용

	path := "/tmp/f4.txt"
	f4, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()

	n, err = io.WriteString(f4, string(buffer))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)

	//----------------------------------------------------------------- 파일에 덧붙이기
	f4, err = os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f4.Close()

	n, err = f4.Write([]byte("hahahahahahahahaha\n"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
