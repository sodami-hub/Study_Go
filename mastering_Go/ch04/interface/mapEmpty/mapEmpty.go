package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var JSONrecord = `{
	"Flag":true,
	"Array":["a","b","c"],
	"Entity": {
		"a1":"b1",
		"a2":"b2",
		"Value":-456,
		"Null":null
	},
	"Meaasge":"Hello Go!"
}`

// 입력으로 들어온 맵의 값들을 구분한다.
// 값들 중 맵이 있따면 typeSwitch를 재귀적으로 호출해 해당 맵의 값들을 조사한다.
func typeSwitch(m map[string]interface{}) {
	for k, v := range m {
		switch c := v.(type) {
		case string:
			fmt.Println("is a string!", k, c)
		case float64:
			fmt.Println("is a float64!", k, c)
		case bool:
			fmt.Println("is a boolean!", k, c)
		case map[string]interface{}:
			fmt.Println("is a map!", k, c)
			typeSwitch(v.(map[string]interface{}))
		default:
			fmt.Printf("...is %v: %T\n", k, c)
		}
	}
	return
}

// for 루프를 이용해 map[string]interface{}의 모든 값을 조사할 수 있다
// exploreMap() 함수는 입력 맵의 내용을 조사한다. 값으로 맵이 있다면 재귀적으로 exploreMap()을 호출한다.
func exploreMap(m map[string]interface{}) {
	for k, v := range m {
		embMap, ok := v.(map[string]interface{})
		// 맵일 경우 한 단계 더 들어간다.
		if ok {
			fmt.Printf("{\"%v\": \n", k)
			exploreMap(embMap)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v : %v\n", k, v)
		}
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("*** Using default JSON record.")
	} else {
		JSONrecord = os.Args[1]
	}

	JSONMap := make(map[string]interface{})
	// json.Unmarshal()은 JSON 데이터를 Go에서 사용하는 값으로 바꾼다. 이값은 보통 구조체이지만 여기서는 map[string]interface{} 맵을 사용한다.
	err := json.Unmarshal([]byte(JSONrecord), &JSONMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	exploreMap(JSONMap)
	fmt.Println()
	typeSwitch(JSONMap)
}
