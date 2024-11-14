/*
구조체의 Tel 필드를 키로 사용한 데이터구조(맵)
insert, delete, search, list 커맨드
*/

package main

import (
	_ "encoding/csv"
	"fmt"
	"os"
	_ "strings"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

const CSVFILE string = "C:\\Users\\leejinhun\\goproject\\mastering_Go\\phoneBook\\pb02\\data.csv"

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Usage: insert|delete|search|list [arguments]")
		return
	}

	//CSVFILE이 존재하지 않는다면 빈 파일을 새로 만든다.
	_, err := os.Stat(CSVFILE)
	// 에러가 nil이 아니라면 파일은 존재하지 않는다.
	if err != nil {
		fmt.Println("Creating", CSVFILE)
		f, err := os.Create(CSVFILE)
		if err != nil {
			f.Close()
			fmt.Println(err)
			return
		}
		f.Close()
	}

	// 파일 정보 가져오기
	fileInfo, err := os.Stat(CSVFILE)

	//정규 파일인가?
	mode := fileInfo.Mode()
	if !mode.IsRegular() { // CSVFILE이 존재해야 되고, 정규 유닉스 파일이어야 된다.
		fmt.Println(CSVFILE, "not a regular file!")
		return
	}
}
