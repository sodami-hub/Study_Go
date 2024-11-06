package main

import (
	"fmt"
	"os"
	"path"
)

type Entry struct {
	Name    string
	Surname string
	Tel     string
}

var data = []Entry{}

func search(key string) *Entry {
	for i, v := range data {
		if v.Surname == key {
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
	data = append(data, Entry{"Mihalis", "Tsoukalos", "2109411236471"})
	data = append(data, Entry{"Mary", "Doe", "21094126871"})
	data = append(data, Entry{"John", "Black", "2109456716123"})
	data = append(data, Entry{"sodam", "lee", "211244416123"})

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
