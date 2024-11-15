/*
구조체의 Tel 필드를 키로 사용한 데이터구조(맵)
insert, delete, search, list 커맨드
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

var data = []Entry{}
var index map[string]int // 전화번호를 키로 해당 데이터(슬라이스)의 인덱스를 가져올 수 있다.
const CSVFILE = "./csv.data"

func listEntry() {
	for _, v := range data {
		fmt.Println("["+v.Name, " : ", v.Surname, " : ", v.Tel, " : ", v.LastAccess+"]")
	}
}

func searchEntry(key string) {
	v, ok := index[key]
	if !ok {
		fmt.Println("없는 전화번호입니다.")
		return
	} else {
		fmt.Println("["+data[v].Name, " : ", data[v].Surname, " : ", data[v].Tel, " : ", data[v].LastAccess+"]")
	}

}

func readCSV() error {

	f, err := os.Open(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, line := range lines {
		if _, ok := index[line[2]]; ok {
			return fmt.Errorf("이미 존재하는 번호입니다.")
		}

		temp := Entry{
			Name:       line[0],
			Surname:    line[1],
			Tel:        line[2],
			LastAccess: line[3],
		}
		data = append(data, temp)
	}
	return nil
}

func saveCSVFile(record []string) error {
	f, err := os.OpenFile(CSVFILE, os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()
	fmt.Println(record)
	csvWriter := csv.NewWriter(f)
	err = csvWriter.Write(record)
	if err != nil {
		fmt.Println(err)
		return err
	}
	csvWriter.Flush()
	return nil
}

func makeIndex() {
	for i := 0; i < len(data); i++ {
		index[data[i].Tel] = i
	}
}

func insertCSV(record []string) error {

	time := time.Now().Unix()
	atNow := strconv.FormatInt(time, 10)
	record = append(record, atNow)

	err := saveCSVFile(record)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("USAGE: list|search|insert|delete [record]")
	}

	err := readCSV()
	if err != nil {
		fmt.Println(err)
		return
	}

	index = make(map[string]int) // 맵 초기화
	makeIndex()

	command := arguments[1]
	switch command {
	case "list":
		listEntry()
	case "search":
		key := arguments[2]
		searchEntry(key)
	case "insert":
		record := arguments[2:]
		if len(record) != 3 {
			fmt.Println("USAGE: insert [Name Surname Tel]")
			return
		}
		tel := strings.ReplaceAll(record[2], "-", "")
		if _, ok := index[tel]; ok {
			fmt.Println("존재하는 전화번호입니다.")
			return
		}
		record[2] = tel
		err = insertCSV(record)
	}
}
