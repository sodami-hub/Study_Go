package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("USAGE: %s URL\n", filepath.Base(os.Args[0]))
		return
	}

	//url.Parse() 함수는 문자열을 URL 구조체로 파싱한다. 따라서 주어진 인수가 유효한 URL이 아니라면 url.Parse()가 알려줄 것이다.
	URL, err := url.Parse(os.Args[1])
	if err != nil {
		fmt.Println("Error in parsing:", err)
		return
	}
	/*
	   코드에서 포인터를 사용하는 보편적인 이유는 뭘까?

	   Go 코드에서 포인터를 사용하는 보편적인 이유는 여러 가지가 있습니다.
	   특히 http.Client와 같은 구조체를 사용할 때 포인터를 사용하는 이유는 다음과 같습니다:
	   1. 메모리 효율성: 구조체를 값으로 전달하면 구조체의 모든 필드가 복사됩니다. 구조체가 클 경우 메모리 사용량이 증가할 수 있습니다.
	   포인터를 사용하면 구조체의 주소만 전달되므로 메모리 사용량이 줄어듭니다.
	   2. 수정 가능성: 포인터를 사용하면 함수나 메서드에서 원본 구조체를 수정할 수 있습니다.
	   값으로 전달하면 복사본이 생성되므로 원본 구조체는 수정되지 않습니다.
	   3. 일관된 인터페이스: 많은 Go 패키지와 라이브러리는 포인터를 사용하여 구조체를 전달하는 것을 기대합니다.
	   예를 들어, http.Client는 포인터로 전달되는 것이 일반적입니다.
	   4. 메서드 호출: 구조체의 메서드가 포인터 리시버를 사용하는 경우, 포인터를 사용하여 메서드를 호출해야 합니다.
	   값으로 전달하면 포인터 리시버 메서드를 호출할 수 없습니다.

	   해당 데이터를 다른 함수로 전달할 것도 아닌데 메모리 사용량이 의미가 있나?
	   그리고 다른 함수로 전달될 것도 아니면 원본 구조체 수정에 대한 요구사항이 없는거 아냐?

	   네, 맞습니다. 만약 http.Client 인스턴스를 다른 함수로 전달하지 않거나, 원본 구조체를 수정할 필요가 없다면, 포인터를 사용하는 것이 반드시 필요한 것은 아닙니다.
	   그러나 여전히 포인터를 사용하는 몇 가지 이유가 있을 수 있습니다:
	   1. 일관된 코드 스타일: Go 코드베이스에서 일관된 스타일을 유지하기 위해 포인터를 사용할 수 있습니다.
	   많은 Go 패키지와 라이브러리는 구조체를 포인터로 전달하는 것을 기대합니다.
	   ✨2. 메서드 호출: http.Client의 메서드가 포인터 리시버를 사용하는 경우, 포인터를 사용해야 합니다.
	   예를 들어, http.Client의 Do 메서드는 포인터 리시버를 사용합니다.
	   ✨3. 기본값 설정: http.Client의 필드 중 일부는 기본값이 설정되어 있습니다.
	   포인터를 사용하면 이러한 기본값을 유지할 수 있습니다.
	*/
	c := &http.Client{
		Timeout: 15 * time.Second,
	}

	// http.NewRequest() 함수는 주어진 메서드, URL, 본문을 갖는 http.Request 객체를 반환한다.
	request, err := http.NewRequest(http.MethodGet, URL.String(), nil)
	if err != nil {
		fmt.Println("GET:", err)
		return
	}

	// http.Do() 함수는 http.Client 를 사용해 HTTP 요청 http.Request를 보낸 뒤 http.Response를 받아 반환한다. http.Do()는 http.Get()을 자세히 표현한 것과 갇다.
	httpData, err := c.Do(request)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	fmt.Println("Status code:", httpData.Status)

	// httputil.DumpResponse() 함수는 서버의 응답을 가져올 때 사용하며 주로 디버깅 용도로 사용한다. 두 번째 인수는 결과의 본문 포함여부를 결정하는 값이다.
	// 여기서는 false로 설정해 본문을 제외한 헤더만 가져온다.
	// 같은 일을 서버 쪽에서 수행한다면 httputil.DumpRequest()를 사용한다.
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Println("header:", string(header))

	// Content-Type의 값을 검색해 응답의 문자열 집합을 알아낸다.
	contentType := httpData.Header.Get("Content-Type")
	characterSet := strings.SplitAfter(contentType, "charset=")
	if len(characterSet) > 1 {
		fmt.Println("Character Set:", characterSet[1])
	}

	// httpData.ContentLength의 값을 읽어 본문의 길이를 구한다.
	if httpData.ContentLength == -1 {
		fmt.Println("ContentLength is unknown!")
	} else {
		fmt.Println("ContentLength:", httpData.ContentLength)
	}

	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
}
