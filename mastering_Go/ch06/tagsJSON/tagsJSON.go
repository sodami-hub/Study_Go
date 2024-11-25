package main

import (
	"encoding/json"
	"fmt"
)

type Password struct {
	Name    string `json:"username"`
	Surname string `json:"surname,omitempty"`
	Year    int    `json:"creationyear,omitempty"`
	Pass    string `json:"-"`
}

func main() {
	var01 := Password{Name: "lee", Surname: "haeden", Pass: "1q2w3e"}

	r, err := json.Marshal(&var01)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// 입력하지 않은 Year 필드는 생략이 됐고, 입력했지만 json을 "-" 처리한 Pass 필드는 json으로 변환되지 않았다.
		fmt.Printf("JSONdata : %s\n", string(r))
	}

	temp := Password{}
	err = json.Unmarshal(r, &temp)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Type: %T : value: %v\n", temp, temp)
	}
}
