/*
에러 랩핑
에러를 감싸서 새로운 에러를 만들어야 할 수도 있다. 예를 들어 파일에서 텍스트를 읽어서 특정 타입의 데이터로 변환하는 경우
파일 읽기에서 발생하는 에러도 필요하지만 텍스트의 몇 번째 줄의 몇 번째 칸에서 에러가 발생했는지도 알면 더 유용하다.
이럴 때 파일 읽기에서 발생한 에러를 감싸고 그 바깥에 줄과 칸 정보를 넣으면 된다.
*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // 1.스캐너 생성
	scanner.Split(bufio.ScanWords)
	//2. 한 단어씩 끊어 읽기
	//Split() 메서드로 Scanner객체를 끊어 읽음 (bufio.ScanWords) -> 한단어씩 끊기
	// (bufio.ScanLines) -> 한 줄씩 끊어 읽음

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(), pos:%d err:%w", pos, err)
		//%w 서식 문자를 통해서 에러 메세지가 중첩돼서 전달된다.
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("failed to readNextInt(),pos:%d err:%w", pos, err)
	}
	return a * b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 반환한다.
// 변환된 숫자, 읽은 글자 수, 에러를 반환한다.
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() { //3. 단어 읽기
		return 0, 0, fmt.Errorf("failed to scan")
	}
	word := scanner.Text()
	number, err := strconv.Atoi(word) //4.문자열을 숫자로 변환
	if err != nil {
		return 0, 0, fmt.Errorf("failed to convert word to int, word:%s err:%w", word, err) //5.에러 감싸기
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) { // 7. 감싸진 에러가 NumError 인지 확인
			fmt.Println("NumberError :", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}

/*
(failed to readNextInt(),pos:4 err:)(failed to convert word to int, word:abc err:)(strconv.Atoi: parsing "abc": invalid syntax
NumberError : strconv.Atoi: parsing "abc": invalid syntax)
*/
