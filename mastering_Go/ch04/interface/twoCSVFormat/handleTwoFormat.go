/*
두 가지 CSV 포맷에서 작동하는 커맨드라인 유틸리티를 구현

서로 다른 CSV 포맷의 레코드는 각각 다른 이름을 가진 Go 구조체에 저장된다. 따라서 모든 CSV 포맷에 대해 sort.Interface를 구현해야 한다.

포맷 1: 이름, 성, 전화번호, 마지막접근시간
포맷 2: 이름, 성, 지역코드, 전화번호, 마지막접근시간
*/

package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const CSV = "./csv.file"

type Entry04 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

type Entry05 struct {
	Name         string
	Surname      string
	locationCode string
	Tel          string
	LastAccess   string
}

var Book01 = []Entry04{}
var Book02 = []Entry05{}

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}
	var firstLine bool = true
	var format bool = true
	for _, v := range lines {
		if firstLine {
			if len(v) == 4 {
				format = true
			} else if len(v) == 5 {
				format = false
			} else {
				return fmt.Errorf("사용하지 않는 형식이다.")
			}
		}
		if format {
			temp := Entry04{
				Name:       v[0],
				Surname:    v[1],
				Tel:        v[2],
				LastAccess: v[3],
			}
			Book01 = append(Book01, temp)
		} else {
			temp := Entry05{
				Name:         v[0],
				Surname:      v[1],
				locationCode: v[2],
				Tel:          v[3],
				LastAccess:   v[4],
			}
			Book02 = append(Book02, temp)
		}
	}
	return nil
}

func main() {
	args := os.Args

	readCSVFile(CSV)
	//sort.Sort(records)

	switch args[1] {
	case "list":
		for _, v := range Book01 {
			fmt.Println("["+v.Name, " : ", v.Surname, " : ", v.Tel, " : ", v.LastAccess+"]")
		}
		for _, v := range Book02 {
			fmt.Println("["+v.Name, " : ", v.Surname, " : ", v.locationCode, " : ", v.Tel, " : ", v.LastAccess+"]")
		}

	}
}

/*
package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sort"
)

// Format 1
type F1 struct {
	Name       string
	Surname    string
	Tel        string
	LastAccess string
}

// Format 2
type F2 struct {
	Name       string
	Surname    string
	Areacode   string
	Tel        string
	LastAccess string
}

type Book1 []F1
type Book2 []F2

// CSVFILE resides in the home directory of the current user
var CSVFILE = ""
var d1 = Book1{}
var d2 = Book2{}

func readCSVFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return err
	}

	var firstLine bool = true
	var format1 = true
	for _, line := range lines {
		if firstLine {
			if len(line) == 4 {
				format1 = true
			} else if len(line) == 5 {
				format1 = false
			} else {
				return errors.New("Unknown File Format!")
			}
			firstLine = false
		}

		if format1 {
			if len(line) == 4 {
				temp := F1{
					Name:       line[0],
					Surname:    line[1],
					Tel:        line[2],
					LastAccess: line[3],
				}
				d1 = append(d1, temp)
			}
		} else {
			if len(line) == 5 {
				temp := F2{
					Name:       line[0],
					Surname:    line[1],
					Areacode:   line[2],
					Tel:        line[3],
					LastAccess: line[4],
				}
				d2 = append(d2, temp)
			}
		}
	}
	return nil
}

// Implement sort.Interface for Book1
func (a Book1) Len() int {
	return len(a)
}

// First based on surname. If they have the same
// surname take into account the name.
func (a Book1) Less(i, j int) bool {
	if a[i].Surname == a[j].Surname {
		return a[i].Name < a[j].Name
	}
	return a[i].Surname < a[j].Surname
}

func (a Book1) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Implement sort.Interface for Book2
func (a Book2) Len() int {
	return len(a)
}

// First based on areacode. If they have the same
// areacode take into account the surname.
func (a Book2) Less(i, j int) bool {
	if a[i].Areacode == a[j].Areacode {
		return a[i].Surname < a[j].Surname
	}
	return a[i].Areacode < a[j].Areacode
}

func (a Book2) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func list(d interface{}) {
	switch T := d.(type) {
	case Book1:
		data := d.(Book1)
		for _, v := range data {
			fmt.Println(v)
		}
	case Book2:
		data := d.(Book2)
		for _, v := range data {
			fmt.Println(v)
		}
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func sortData(data interface{}) {
	// type switch
	switch T := data.(type) {
	case Book1:
		d := data.(Book1)
		sort.Sort(Book1(d))
		list(d)
	case Book2:
		d := data.(Book2)
		sort.Sort(Book2(d))
		list(d)
	default:
		fmt.Printf("Not supported type: %T\n", T)
	}
}

func main() {
	if len(os.Args) != 1 {
		CSVFILE = os.Args[1]
	} else {
		fmt.Println("No data file!")
		return
	}

	_, err := os.Stat(CSVFILE)
	// If the CSVFILE does not exist, terminate the program
	if err != nil {
		fmt.Println(CSVFILE, "does not exist!")
		return
	}

	fileInfo, err := os.Stat(CSVFILE)
	// Is it a regular file?
	mode := fileInfo.Mode()
	if !mode.IsRegular() {
		fmt.Println(CSVFILE, "not a regular file!")
		return
	}

	err = readCSVFile(CSVFILE)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(d1) != 0 {
		sortData(d1)
	} else {
		sortData(d2)
	}
}
*/
