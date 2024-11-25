package main

import (
	"encoding/xml"
	"fmt"
)

type Employee struct {
	XMLName   xml.Name `xml:"employee`
	ID        int      `xml:"id,attr"`
	FirstName string   `xml:"name>first"`
	LastName  string   `xml:"name>last"`
	Height    float32  `xml:"height,omitempty"`
	Address
	Comment string `xml:",comment"`
}

/*
XML 데이터의 구조체를 정의했다. 그러나 여기서는 이름과 타입에 관한 추가적인 정보가 있다.
XMLName 필드는 XML 레코드의 이름이고 여기서는 employee가 된다. ",comment" 태그를 갖는 필드는 XML의 주석이며 출력에서 따로 포맷팅한다.
attr태그가 있는 필드는 출력에서 해당 필드명(여기서는 id)이 속성으로 나타난다.
name>first 표현은 Go가 first 태그를 name 태그 내부에 임베딩하라는 의미이다.
*/

type Address struct {
	City, Country string
}

func main() {
	r := Employee{ID: 7, FirstName: "Lee", LastName: "sodam"}
	r.Comment = "loved, liked, hahaha"
	r.Address = Address{City: "seongnam", Country: "ROK"}

	output, err := xml.MarshalIndent(&r, " ", " ")

	if err != nil {
		fmt.Println(err)
		return
	} else {
		output = []byte(xml.Header + string(output))
		fmt.Println(string(output))
	}
}
