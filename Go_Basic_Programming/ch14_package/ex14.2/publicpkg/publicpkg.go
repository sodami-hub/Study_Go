package publicpkg

import "fmt"

const (
	PI = 3.1415   // 1. 공개되는 상수 - 대문자로 시작하는 패키지 전역 상수는 외부로 공개
	pi = 3.141516 // 2. 공개되지 않는 상수 - 소문자로 시작하는 패키지 전역 상수는 비공개
)

var ScreenSize int = 1080 // 100. 공개되는 변수
var screenHeight int      // 공개되지 않는 변수

func PublicFunc() { // 공개되는 함수
	const MyConst = 100 // 3. 공개되지 않는다. - 상수는 대문자로 시작하더라도 하무 내부에서 선언되면 패키지 외부로 공개되지 않는다.
	fmt.Println("This is a public function", MyConst)
}

func privateFunc() { // 비공개 함수
	fmt.Println("this is a private function")
}

type MyInt int       // 공개되는 별칭 타입
type myString string // 비공개 별칭 타입

type MyStruct struct { // 공개되는 구조체
	Age  int    // 4. 공개 필드 - 공개되는 구조체의 대문자로 시작하는 필드는 외부로 공개
	name string // 5. 비공개 필드 - 공개되는 구조체라도 소문자로 시작되는 필드는 외부로 비공개
}

func (m MyStruct) PublicMethod() { // 6. 공개되는 메서드 - 공개되는 구조체의 공개되는 메소드는 외부로 공개
	fmt.Println("this is public method")
}

func (m MyStruct) privateMethod() { // 비공개 메서드
	fmt.Println("this is private method")
}

type myPrivateStruct struct { // 공개되지 않는 구조체
	Age  int    // 7. 비공개 구조체 필드 - 대문자로 시작하지만 속한 구조체가 비공개이므로 외부에 공개되지 않는다.
	name string // 비공개 구조체 필드
}

func (m myPrivateStruct) PrivateMethod() { // 비공개 메소드
	fmt.Println("this is a private method")
}
