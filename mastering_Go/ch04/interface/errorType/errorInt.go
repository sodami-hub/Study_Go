package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type emptyFile struct {
	Ended bool
	Read  int
}

// error 인터페이스 구현
func (e emptyFile) Error() string {
	return fmt.Sprintf("Ended with io.EOF (%t) but read (%d) bytes", e.Ended, e.Read)
}

// 파일이 빈 파일인지 아닌지 체크하는 함수.
func isFileEmpty(e error) bool {
	v, ok := e.(emptyFile)
	if ok {
		if v.Read == 0 && v.Ended == true { // 0바이트를 읽었고, 파일의 끝에 도달했다면 해당 파일은 빈 파일이다.
			return true
		}
	}
	return false
}

func readFile(file string) error {
	var err error
	fd, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fd.Close()

	reader := bufio.NewReader(fd)
	n := 0
	for {
		line, err := reader.ReadString('\n') // 개행문자가 나올 때 까지 읽어들인다.
		n += len(line)
		if err == io.EOF {
			if n == 0 {
				return emptyFile{true, n}
			}
			break
		} else if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("usage : errorInt <file1> [<file2>...]")
		return
	}

	for _, file := range flag.Args() {
		err := readFile(file)
		if isFileEmpty(err) {
			fmt.Println(file, err)
		} else if err != nil {
			fmt.Println(file, err)
		} else {
			fmt.Println(file, "is OK.")
		}
	}
}
