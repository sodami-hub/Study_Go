package main

import (
	"fmt"
	"goproject/usepkg/custompkg"

	"github.com/guptarohit/asciigraph"
	"github.com/tuckersGo/musthaveGo2/ch14/expkg"
)

func main() {
	custompkg.PrintCustom()
	expkg.PrintSample()

	data := []float64{3, 4, 5, 6, 9, 7, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
}

/*
1. goproject/usepkg> go mod init goproject/usepkg
2. custompkg.go 와 usepkg.go 파일을 생성하고 코드 작성
3. goproject/usepkg> go mod tidy    -> 외부 라이브러리를 다운받고 go.mod, go.sum 파일에 외부 라이브러리에대한 정보를 작성한다.
4. goproject/usepkg> go build -> 실행파일 생성
5. goproject/usepkg> ./usepkg.exe

*/
