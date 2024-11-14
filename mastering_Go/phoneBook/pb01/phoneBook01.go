/*
random() 을 사용해서 랜덤한 이름과 성, 전화번호를 100개 생성해서 전화번호부를 만들고
리스트확인 및 검색을 제공
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}
var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(len int64) string {
	temp := ""
	startChar := "!"
	var i int64 = 1
	for {
		myRand := random(32, 58)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == len {
			break
		}
		i++
	}
	return temp
}

func populate(n int) {
	for i := 0; i < n; i++ {
		name := getString(4)
		surname := getString(5)
		n := strconv.Itoa(random(100, 199))
		data = append(data, Entry{name, surname, n})
	}
}

func search(key string) *Entry {
	for i, v := range data {
		if v.Tel == key {
			return &data[i]
		}
	}
	return nil
}

func list() {
	for _, v := range data {
		fmt.Println(v)
	}
}

func main() {
	args := os.Args
	if len(args) == 1 {
		exe := path.Base(args[0])
		fmt.Printf("Usage: %s search | list <arguments>\n", exe)
		return
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)

	n := 100
	populate(n)
	fmt.Printf("Data has %d entries.\n", len(data))

	switch args[1] {
	case "search":
		if len(args) != 3 {
			fmt.Println("Usage : search Surname")
			return
		}
		result := search(args[2])
		if result == nil {
			fmt.Println("entry not fount : ", args[2])
			return
		}
		fmt.Println(*result)
	case "list":
		list()
	default:
		fmt.Println("not a valid option")
	}
}
