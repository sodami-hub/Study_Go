package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	LOGFILE := path.Join(os.TempDir(), "mGo1.log")
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	LstdFlags := log.Ldate | log.Ltime
	iLog := log.New(f, "LNUM", LstdFlags)
	iLog.Println("loglogloglog") //LNUM2024/11/06 13:28:21 loglogloglog

	iLog.SetFlags(log.Lshortfile | log.LstdFlags) // log.Lshorfile -> 로그 항목을 출력하는 소스코드의 파일명과 줄 번호를 추가
	iLog.Println("another log entry!")            //customLogLineNumber.go:26: another log entry!
	iLog.SetFlags(log.Llongfile | log.LstdFlags)  // log.Llongfile -> 로그 항목을 출력하는 소스코드의 파일명을 풀패스로 넣고, 줄 번호를 추가
	iLog.Println("another log longlong entry!")   // /home/sodami/Documents/Work_Space/goproject/mastering_Go/ch01/customLogLineNumber/customLogLineNumber.go:28: another log longlong entry!
}
