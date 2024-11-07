package main

import (
	"fmt"
	"os"            //운영체제와 상호작용
	"path/filepath" // 문자열로 읽기에는 너무 많은 디렉터리를 가지고 있는 PATH 변수를 다루기 위한 패키지
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Print("인수를 넣으세요.")
	}

	file := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		// 파일이 존재하는가?
		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()
			//일반 파일인가?
			if mode.IsRegular() {
				//실행 파일인가?
				if mode&0111 != 0 {
					fmt.Println(fullPath)
					return
				}
			}
		}
	}
}
