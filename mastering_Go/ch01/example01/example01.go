// which.go를 수정해서 가능한 모든 실행 파일을 찾을 수 있게 만들어보자.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("USAGE: 1개의 인수를 넣어주세요.")
		return
	}

	fileName := args[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, fileName)

		fileInfo, err := os.Stat(fullPath)
		if err == nil {
			mode := fileInfo.Mode()

			if mode.IsRegular() {
				if mode&0111 != 0 {
					fmt.Println(fullPath)
				}
			}
		}
	}

}
