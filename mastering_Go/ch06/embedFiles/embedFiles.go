package main

import (
	_ "embed" // 임베딩하려면 embed 패키지가 필요하다. 직접 사용하지는 않으므로 _ 를 이용해서 에러가 나지 않도록 한다.
	"fmt"
	"os"
)

// 줄을 //go:embed로 시작하면 Go는 이를 특별한 주석으로 취급한다. 주석의 뒤쪽에 적힌 경로에 해당하는 파일을 임베딩한다.

//go:embed gopher.webp
var f1 []byte

//go:embed testfile.txt
var f2 string

func writeToFile(s []byte, path string) error {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer fd.Close()

	n, err := fd.Write(s)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n)
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("more args")
		return
	}

	fmt.Println("f1:", len(f1), "f2:", len(f2))

	switch arguments[1] {
	case "1":
		filename := "/tmp/temporary.webp"
		err := writeToFile(f1, filename)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "2":
		fmt.Print(f2)
	default:
		fmt.Println("not a valid option!")
	}
}
