package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

// 보통 파일에서 데이터를 읽어오지만 여기서는 아래 변수를 사용할 예정이다.
var yamlfile = `
image: Golang
matrix:
  docker: python
  version: [2.7, 3.9]
`

// 구조체 Mat에 YAML 파일에 관련된 두 가지의 필드를 정의한다.
// Version 필드는 float32 슬라이스다. 이름을 정의하지 않았기 때문에 version이 될 것이다.
// flow 키워드는 마샬링을 할 때 플로 스타일을 사용하도록 지정하고 이는 구조체, 시퀀스, 맵에서 유용하게 사용된다.
type Mat struct {
	DockerImage string    `yaml:"docker"`
	Version     []float32 `yaml:",flow"`
}

// YAML 구조체는 Mat 구조체를 임베딩하고 YAML 파일의 image와 대응하는 Image 필드를 가지고 있다.
type YAML struct {
	Image  string
	Matrix Mat
}

func main() {
	data := YAML{}

	err := yaml.Unmarshal([]byte(yamlfile), &data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After Unmarshal (Structure):\n%v\n\n", data)

	d, err := yaml.Marshal(&data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("After Marshal (YAML code):\n%s\n", string(d))
}
