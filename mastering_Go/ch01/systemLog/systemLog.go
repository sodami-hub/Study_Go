package main

import (
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.go")

	if err != nil {
		log.Println(err)
		return
	} else {
		log.SetOutput(sysLog) // SetOutpu() 을 호출해서 모든 로그 정보를 sysLog 변수에 전달돼고
		// 전달된 로그는 syslog.LOG_SYSLOG에 로그를 남긴다. syslog.New의 두번째 인수도 함께 남긴다.
		log.Print("Everything is fine!")
		log.Print("Really?")
	}
}

// 남겨진 로그는 맥OS는 /var/log/system.log 에서 찾을 수 있으며
// 리눅스 환경에서는 journalctl -xe 를 실행해서 확인할 수 있다.
// 또는 cat /var/log/syslog |grep Really? 이런식으로도 검색할 수 있다.
