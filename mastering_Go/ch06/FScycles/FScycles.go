/*
유닉스 시스템의 파일 순환 참조 해결방법
- 방문한 모든 데릭터리 경로를 맵에 기록하고 해당 경로가 두 번 이상 나타난다면 사이클이 있다고 판단한다.
그 맵은 visited라고 부르고 map[string]int 로 정의한다.

*/

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var visited = map[string]int{}

func walkFunction(path string, info os.FileInfo, err error) error {

	// 먼저 존재하는 경로인지 확인한다.(os.Stat())
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil
	}

	// os.Lstat()을 사용해서 심볼릭 링크를 처리한다.  os.Lstat()은 os.Stat()과 달리 심볼릭 링크를 따라가지 않고 정보를 반환한다.
	fileInfo, _ = os.Lstat(path)
	mode := fileInfo.Mode()

	// 정규 디렉터리인지 확인한다.
	if mode.IsDir() {
		abs, _ := filepath.Abs(path)
		_, ok := visited[abs] // 맵에 abs 값이 있으면(이미 방문한 경로이면) ok는 true와 _는 value를 반환한다.
		if ok {
			fmt.Println("Found Cycle", abs)
			return nil
		}
		visited[abs]++
		return nil
	}

	// 심볼릭 링크가 디렉터리를 가리키고 있는지 찾는다.
	// fileInfo가 심볼릭 링크인지 확인한다.
	// fileInfo의 파일모드와 심볼릭 링크의 비트 마스크(os.ModeSymlink)를 & 연산해서
	// true이면 ...
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		temp, err := os.Readlink(path)
		if err != nil {
			fmt.Println("os.Readlink():", err)
			return err
		}
		// EvalSymlinks() 함수는 심볼릭 링크가 가리키고 있는 곳을 찾아내려고 사용한다.
		newPath, err := filepath.EvalSymlinks(temp)
		if err != nil {
			return nil
		}
		linkFileInfo, err := os.Stat(newPath)
		if err != nil {
			return err
		}

		linkMode := linkFileInfo.Mode()
		if linkMode.IsDir() {
			fmt.Println("Following...", path, "--->", newPath)
			abs, _ := filepath.Abs(newPath)
			_, ok := visited[abs]
			if ok {
				fmt.Println("Found cycle!", abs)
				return nil
			}
			visited[abs]++
			err = filepath.Walk(newPath, walkFunction)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("not enough arguments!")
		return
	}

	Path := arguments[1]
	// filepath.Walk() 함수는 지정된 경로(첫번째 인수)를 순회하면서 각 파일이나 디렉토리에 대해 지정된 함수(두번째 인수)를 호출한다.
	// 이 함수는 파일 시스템을 탐색하고, 각 파일이나 디렉토리에 대해 특정 작업을 수행할 때 유용하다.
	err := filepath.Walk(Path, walkFunction)
	if err != nil {
		fmt.Println(err)
	}

	for k, v := range visited {
		if v > 1 {
			fmt.Print(k, v)
		}
	}

}
