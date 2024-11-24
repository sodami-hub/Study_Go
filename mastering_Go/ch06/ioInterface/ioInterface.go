package main

import (
	"bufio"
	"fmt"
	"io"
)

type S1 struct {
	F1 int
	F2 string
}

type S2 struct {
	F1   S1
	text []byte
}

// S1의 io.Reader 인터페이스 구현
// S1의 포인터를 사용해 메서드에서 변경 사항이 구조체에 저장될 수 있도록한다.
func (s *S1) Read(p []byte) (n int, err error) {
	fmt.Print("Give me your name : ")
	fmt.Scanln(&p)
	fmt.Println("buf size", len(p), cap(p))
	s.F2 = string(p)
	return len(p), nil
}

// S1의 io.Writer 인터페이스 구현
func (s *S1) Write(p []byte) (n int, err error) {
	fmt.Println("n ->", n) // 반환값에 선언한 변수는 자동으로 초기화가 된다.
	if s.F1 < 0 {
		return -1, nil
	}

	for i := 0; i < s.F1; i++ {
		fmt.Printf("%s", p)
	}

	fmt.Println()
	return s.F1, nil
}

// 아래 두 함수는 표준 라이브러리의 bytes.Buffer.ReadByte의 구현이다.
func (s S2) eof() bool {
	return len(s.text) == 0
}
func (s *S2) readByte() byte {
	// 여기서는 eof() 함수에서 체크를 이미 했다고 가정한다.
	temp := s.text[0]
	s.text = s.text[1:]
	return temp
}

// 모든 데이터를 읽으면 관련 구조체의 필드가 비워진다(readByte()).
func (s *S2) Read(p []byte) (n int, err error) {
	if s.eof() {
		err = io.EOF
		return
	}

	l := len(p)
	if l > 0 {
		for n < l {
			p[n] = s.readByte()
			n++
			if s.eof() {
				s.text = s.text[0:0]
				break
			}
		}
	}
	return
}

func main() {
	s1var := S1{4, "Hello"}
	fmt.Println(s1var)

	buf := make([]byte, 2)
	fmt.Println("buf size", len(buf), cap(buf))
	_, err := s1var.Read(buf)
	fmt.Println("buf size", len(buf), cap(buf))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Read", s1var.F2)

	_, _ = s1var.Write([]byte("Hello There!"))

	s2var := S2{F1: s1var, text: []byte("Hello world!")}

	r := bufio.NewReader(&s2var)

	for {
		n, err := r.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("*", err)
			break
		}
		fmt.Println("**", n, string(buf[:n]))
	}
}
