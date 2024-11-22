// mypackage 패키지를 를 임포트해서 실행한다.

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sodami-hub/Study_Go/mastering_Go/ch05/mypackage"
)

var MIN = 0
var MAX = 26

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func getString(length int64) string {
	startChar := "A"
	temp := ""
	var i int64 = 1

	for {
		myRand := random(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		temp = temp + newChar
		if i == length {
			break
		}
		i++
	}
	return temp
}

func main() {
	mypackage.Hostname = "localhost"
	mypackage.Username = "root"
	mypackage.Password = "admin"
	mypackage.Database = "go"
	mypackage.Port = 3306

	data, err := mypackage.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}

	SEED := time.Now().Unix()
	rand.Seed(SEED)
	random_username := getString(5)

	t := mypackage.Userdata{
		Username:   random_username,
		Name:       "sodam",
		Surname:    "lee",
		Descrption: "놀자",
	}

	id := mypackage.AddUser(t)
	if id == -1 {
		fmt.Println("adduser error", t.Username)
	}
}
