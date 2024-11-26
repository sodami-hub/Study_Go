package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetSize(path string) (int64, error) {
	contents, err := os.ReadDir(path)
	if err != nil {
		return -1, err
	}

	var total int64
	for _, entry := range contents {
		// 디렉터리 항목을 방문한다.
		if entry.IsDir() {
			// 디렉터리를 처리할 때는 계속 깊숙하게 들어가야 한다.
			temp, err := GetSize(filepath.Join(path, entry.Name()))
			if err != nil {
				return -1, err
			}
			total += temp
			// 디렉터리가 아닌 항목들 각각의 크기를 구한다.
		} else {
			info, err := entry.Info()
			if err != nil {
				return -1, err
			}
			// int64 값을 반환한다.
			total += info.Size()
		}
	}
	return total, nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need a <Directory>")
		return
	}

	root, err := filepath.EvalSymlinks(arguments[1])
	fileInfo, err := os.Stat(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	fileInfo, _ = os.Lstat(root)
	mode := fileInfo.Mode()
	if !mode.IsDir() {
		fmt.Println(root, "not a directory!")
		return
	}

	i, err := GetSize(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Total Size:", i)
}
