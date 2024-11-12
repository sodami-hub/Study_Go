// 입력에 대한 검사를 수행하기 전에 전체 레코드를 읽고 분할할 때 다른 방식을 사용한다.
// 또한 처리하고자 하는 레코드의 필드가 정확한 개수인지도 검사한다.
// 각각의 레코드는 이름, 성, 전화번호까지 총 3개의 필드를 갖고 있어야 한다.

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func matchNameSur(s string) bool {
	t := []byte(s)
	// ^, $ 줄의 시작과 끝, 대문자로 시작, 그다음부터는 아무 길이(*)의 소문자로 된 문자열
	re := regexp.MustCompile(`^[A-Z][a-z]*$`)
	return re.Match(t)
}

func matchTel(s string) bool {
	t := []byte(s)
	// +,- 로 시작(생략가능 '?'), 1개 이상의 아무길이(+)의 숫자(\d)로 이루어진 문자열
	re := regexp.MustCompile(`^[+-]?\d+$`)
	return re.Match(t)
}

func matchRecord(s string) bool {
	fields := strings.Split(s, ",")
	if len(fields) != 3 {
		return false
	}
	if !matchNameSur(fields[0]) {
		return false
	}
	if !matchNameSur(fields[1]) {
		return false
	}
	return matchTel(fields[2])
}

func main() {
	args := os.Args

	res := matchRecord(args[1])
	fmt.Println(res)
}
