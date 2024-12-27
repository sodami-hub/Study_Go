package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func main() {
	reqBody := new(bytes.Buffer)
	w := multipart.NewWriter(reqBody)

	for k, v := range map[string]string{
		"data": time.Now().Format(time.RFC3339),
		"desc": "Form values with attached files",
	} {
		err := w.WriteField(k, v)
		if err != nil {
			fmt.Println(err)
		}
	}

	/* 위의 for문은 아래의 것과 같은 내용이다.
	myMap := map[string]string{
		"a": "bcd",
		"b": "cde",
	}
	for k,v := range myMap {
		err := w.WriteField(k,v)
		if err != nil {
			fmt.Print(err)
		}
	}
	*/

	for i, file := range []string{
		"./files/hello.txt",
		"./files/goodbye.txt",
	} {
		// CreateFormField 메서드는 매개변수로 필드명과 파일명을 받는다. 서버는
		// 받은 파일명을 이용하여 MIME파트를 파싱한다. 첨부파일명과 일치할 필요는 없다.
		filePart, err := w.CreateFormFile(fmt.Sprintf("file%d", i+1),
			filepath.Base(file))

		if err != nil {
			fmt.Println(err)
			return
		}
		f, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			return
		}

		// 파일을 열고 파일의 내용을 MIME 파트의 writer로 복사한다.
		_, err = io.Copy(filePart, f)
		_ = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// 요청 보디에 파트 추가가 완료되면 반드시 멀티파트 writer를 닥아야 요청 보디가 바운더리를 추가하는 작업을 올바르게 마무리한다.
	err := w.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, "https://httpbin.org/post", reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer func() { _ = resp.Body.Close() }()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("expected status %d: actual status %d", http.StatusOK, resp.StatusCode)
		return
	}

	fmt.Printf("\n%s\n", b)
}
