package main

import (
	"encoding/json"
	"fmt"
)

// 편의상 하드코딩한 데이터를 활용한다.
type UseAll struct {
	Name    string `json:"username"`
	Surname string `json:"surname`
	Year    int    `json:"created"`
}

// 위 데이터가 알려주는 사실은 UseAll 구조체의 Name 필드가 JSON 레코드에서는 username으로 변환한다는 사실이다.
// 이 정보들은 JSON 데이터를 마샬링 및 언마샬링할 때 사용한다.
// 이를 제외하고는 구조체 UseAll은 일반 Go 구조체와 동일하게 사용할 수 있다.

func main() {
	useall := UseAll{Name: "Mike", Surname: "Lee", Year: 2019}

	// 일반 구조체를 JSON 데이터로 인코딩하면 -> 구조체의 필드를 가지고 있는 JSON 레코드로 변환
	t, err := json.Marshal(&useall)
	// json.Marshal() 함수는 구조체 변수의 포인터를 매개변수로 받는다. 인코딩한 정보가 담긴 바이트 슬라이스와 에러를 반환한다.
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Value %s\n", t)
	}

	// 주어진 JSON데이터를 문자열로 디코딩한다.
	str := `{"username":"M.", "surname":"TS","created":2019}`
	// 문자열을 바이트 슬라이스로 변환한다.
	jsonRecord := []byte(str)

	// 결과를 저장할 구조체 변수 생성
	temp := UseAll{}
	// JSON데이터의 바이트 슬라이스를 구조체로 변환한다.
	err = json.Unmarshal(jsonRecord, &temp)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Data type : %T with value : %v\n", temp, temp)
	}
}
