package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type User struct {
	First string
	Last  string
}

func main() {

	user := User{First: "lee", Last: "sodam"}

	userData, _ := json.Marshal(user)
	u := bytes.NewReader(userData)

	req, err := http.NewRequest(http.MethodPost, "http://localhost:1234/postuser", u)
	if err != nil {
		fmt.Println("error in req", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	resp, err := c.Do(req)
	if err != nil {
		fmt.Println("resp error", err)
		return
	}
	defer resp.Body.Close()

	if resp == nil {
		fmt.Println(http.StatusNotFound)
		return
	}
	fmt.Println("from server status : ", resp.StatusCode)
}
