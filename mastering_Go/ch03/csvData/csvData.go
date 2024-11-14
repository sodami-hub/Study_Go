/*
> go run csvData.go ./csv.data ./data.csv
디렉터리의 csv.data를 읽어서(readCSVFile)(첫 번째 실행 인수) data.csv에 출력한다.(saveCSVFile)(두번째 실행인수)
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Record struct {
	Name       string
	Surname    string
	Number     string
	LastAccess string
}

var myData = []Record{}

// 첫번째 실행인수인 data의 파일 경로를 파라미터로 받는다.
func readCSVFile(filepath string) ([][]string, error) {
	_, err := os.Stat(filepath) // 파일 경로의 존재 여부와 정규 파일인지에 대한 체크
	if err != nil {
		return nil, err
	}

	f, err := os.Open(filepath) // 경로의 파일을 연다
	if err != nil {
		return nil, err
	}

	defer f.Close()

	// CSV 파일을 한번에 읽어온다.
	// csv.NewReader()는 각 줄의 필드들을 분리해주기 때문에 2차원 슬라이스가 필요하다. lines의 타입은 [][]string이다.
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, err
}

func saveCSVFile(filepath string) error {
	csvfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer csvfile.Close()

	csvwriter := csv.NewWriter(csvfile)
	// 기본 구분자를 탭으로 바꾼다.
	csvwriter.Comma = '\t'
	for _, row := range myData {
		temp := []string{row.Name, row.Surname, row.Number, row.LastAccess}
		_ = csvwriter.Write(temp)
	}
	csvwriter.Flush()
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("csvData input output!")
		return
	}

	input := os.Args[1]
	output := os.Args[2]
	lines, err := readCSVFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 첫번째 실행인수의 파일을 열어서 이차원 슬라이스로 받고
	// 슬라이스를 다시 구조체의 슬라이스로 변환한다
	for _, line := range lines {
		temp := Record{ // 읽어온 데이터를 구조체로 바꾸고
			Name:       line[0],
			Surname:    line[1],
			Number:     line[2],
			LastAccess: line[3],
		}
		myData = append(myData, temp) // 구조체들을 슬라이스(전역변수)로 정리
		fmt.Println(temp)
	}

	err = saveCSVFile(output) // 두번째 실행인수의 경로에 csv 파일을 저장
	if err != nil {
		fmt.Println(err)
		return
	}
}
