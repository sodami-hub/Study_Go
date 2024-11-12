package main

import "fmt"

type Entry struct {
	Name    string
	Surname string
	Year    int
}

// Go에 의해서 초기화
func zeroS() Entry {
	return Entry{}
}

// 사용자에 의해 초기화
func initS(N, S string, Y int) Entry {
	if Y < 2000 {
		return Entry{Name: N, Surname: S, Year: 2000}
	} else {
		return Entry{Name: N, Surname: S, Year: Y}
	}
}

// Go에 의해서 초기화하고 포인터 반환
func zeroPtoS() *Entry {
	t := &Entry{}
	return t
}

// 사용자가 초기화하고 포인터 반환
func initPtoS(N, S string, Y int) *Entry {
	if len(S) == 0 {
		return &Entry{Name: N, Surname: "unknown", Year: Y}
	}
	return &Entry{Name: N, Surname: S, Year: Y}
}

func main() {
	s1 := zeroS()
	p1 := zeroPtoS()
	fmt.Println("s1 :", s1, " p1 :", *p1) // 문자열의 제로값은 빈 문자열이다.
	s2 := initS("sodam", "lee", 2019)
	p2 := initPtoS("jinhun", "lee", 1984)
	fmt.Println("s1 :", s2, " p1 :", *p2)

	fmt.Println("Year :", s1.Year, s2.Year, p1.Year, p2.Year)

	pS := new(Entry)
	fmt.Println("pS:", pS)
}
