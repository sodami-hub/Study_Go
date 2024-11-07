// which.go를 개선해서 여러개의 실행 바이너리를 검색할 수 있도록 바꾼다.

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("1개 이상의 실행 인수를 입력하쇼")
		return
	}

	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for i := 1; i < len(args); i++ {
		fileName := args[i]

		for _, dir := range pathSplit {
			fullPath := filepath.Join(dir, fileName)

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

}
