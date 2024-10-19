package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFile(filename string) (string, error) {
	file, err := os.Open(filename) //1. 파일 열기
	if err != nil {
		return "", err //2. 에러발생하면 에러 반환
	}
	defer file.Close()             //3. 함수 종료 직전 파일 닫기
	rd := bufio.NewReader(file)    //4. 파일 내용 읽기
	line, _ := rd.ReadString('\n') // '\n'이 나올 때 까지 파일 읽기.->한 줄 읽기
	return line, nil
}

func WriteFile(filename string, line string) error {
	file, err := os.Create(filename) //5. 파일 생성
	if err != nil {                  //6. 에러 나면 에러 반환
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintln(file, line) //7. 파일에 문자열 쓰기
	return err
}

const filename string = "data.txt"

func main() {
	line, err := ReadFile(filename) //8. 파일 읽기 시도
	if err != nil {
		err = WriteFile(filename, "This is WriteFile") //9. 파일 생성
		if err != nil {                                //10. 9번의 파일 생성 과정에서의 에러 처리
			fmt.Println("파일 생성에 실패", err)
			return
		}
		line, err = ReadFile(filename) //11. 9번에서 생성된 파일 다시 읽기 시도
		if err != nil {
			fmt.Println("파일 읽기에 실패", err)
			return
		}
	}

	fmt.Println("파일 내용 : ", line)
}
