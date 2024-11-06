package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	iLog := log.New(f, "iLog", log.LstdFlags) // iLog 변수를 사용해서 io.Writer(f)에 로그를 남긴다. "iLog"와 flag가 프리픽스로 붙는다.
	iLog.Println("Hello there")
	iLog.Println("Mastering Go 3rd Edition!")
}
