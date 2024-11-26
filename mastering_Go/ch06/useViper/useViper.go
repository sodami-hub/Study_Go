package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// 플래그에 대한 추가적인 앨리어스를 만들고자 사용. --password 플래그 대신 --pass, --ps를 사용할 수 있다.
func aliasNormalizeFunc(f *pflag.FlagSet, n string) pflag.NormalizedName {
	switch n {
	case "pass":
		n = "password"
	case "ps":
		n = "password"
	}
	return pflag.NormalizedName(n)
}

func main() {
	// name이라는 새 플래그를 만들고 -n 으로 접근할 수 있게 했다. 기본값은 MIKE이며 설명은 Name parameter 이다.
	pflag.StringP("name", "n", "MIKE", "Name parameter")

	// 같은 방식으로 password 플래그를 만들고, 표준화된 함수를 등록해 password 플래그의 앨리어스를 만들었다.
	pflag.StringP("password", "p", "hardToGuess", "Password")
	pflag.CommandLine.SetNormalizeFunc(aliasNormalizeFunc)

	// pflag.Parse()는 모든 커맨드라인 플래그를 정의한 뒤 호출해야 된다. 이는 커맨드라인 플래그를 파싱해 정의한 플래그의 값으로 넣는다.
	pflag.Parse()
	// viper.BindPFlags()를 호출하면 모든 플래그를 viper 패키지에 사용할 수 있게 된다.
	// 다시말해 pflag 플래그들(pflag.FlagSet) 을 viper에 바인등한다.
	viper.BindPFlags(pflag.CommandLine)

	// 플래그를 읽는다.
	name := viper.GetString("name")
	password := viper.GetString("password")

	fmt.Println(name, password)

	// 환경 변수를 읽는다.
	viper.BindEnv("PATH")
	val := viper.Get("PATH")
	if val != nil {
		fmt.Println("PATH:", val)
	}

	// 환경 변수를 설정한다.
	viper.Set("GOMAXPROCS", 16)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)
}
