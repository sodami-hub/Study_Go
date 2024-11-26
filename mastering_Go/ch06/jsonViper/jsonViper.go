package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type ConfigStructure struct {
	MacPass    string `mapstructure:"macos"`
	LinuxPass  string `mapstructure:"linux"`
	WindowPass string `mapstructure:"windows"`
	PostHost   string `mapstructure:"postgres"`
	MySQLHost  string `mapstructure:"mysql"`
	MongoHost  string `mapstructure:"mongodb"`
}

// 중요한 점은 JSON 파일에 설정을 저장하더라도 구조체는 JSON 설정 파일의 필드로 json 대신 mapstructure를 사용한다는 것이다.

var CONFIG = ".config.json"

func main() {
	if len(os.Args) == 1 {
		fmt.Println("use defaultfile", CONFIG)
	} else {
		CONFIG = os.Args[1]
	}

	// json파일을 사용한다고 선언
	viper.SetConfigType("json")
	// 설정파일의 경로
	viper.SetConfigFile(CONFIG)
	// 설정파일의 내용을 출력한뒤 설정 파일을 읽고 파싱한다.
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	// 바이퍼 패키지는 설정 파일이 실제로 존재하고 읽을 수 있는지 체크하지 않는다. 파일을 찾을 수 없거나 읽을 수 없다면
	// 아래 함수는 빈 설정 파일을 읽는 것처럼 동작할 것이다.
	viper.ReadInConfig()

	//viper.IsSet() 을 호출하면 macos라는 키가 설정에 있는지 체크한다. 있다면 viper.Get("macos")를 이용해 값을 읽어 화면에 출력한다.
	if viper.IsSet("macos") {
		fmt.Println("macos:", viper.Get("macos"))
	} else {
		fmt.Println("macos not set!")
	}

	// 먼저 active 키가 있는지 체크한다. 값이 true라면 세가지 키의 값을 더 읽는다.
	if viper.IsSet("active") {
		// active 키는 불리언 값을 갖고 있기 때문에 viper.GetBool()로 값을 읽는다.
		value := viper.GetBool("active")
		if value {
			postgres := viper.Get("postgres")
			mysql := viper.Get("mysql")
			mongo := viper.Get("mongodb")
			fmt.Println("P:", postgres, "My:", mysql, "Mo:", mongo)
		}
	} else {
		fmt.Println("active is not set")
	}

	// 없는 키를 읽으면 실패한다.
	if !viper.IsSet("DoesNotExist") {
		fmt.Println("DoesNotExist is not set")
	}

	var t ConfigStructure
	err := viper.Unmarshal(&t)
	if err != nil {
		fmt.Println(err)
		return
	}
	PrettyPrint(t)
}

func PrettyPrint(v interface{}) (err error) {
	// json.MarshalIndent => Marshal + Indent
	b, err := json.MarshalIndent(v, "", "\t")
	if err == nil {
		fmt.Println(string(b))
	}
	return err
}
