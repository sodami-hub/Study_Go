package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRestrictPrefix(t *testing.T) {
	/*
		http.FileServer에서 서빙되는 정적 파일들의 디렉터리는 현재 디렉터리의 부모디렉터리안에 존재한다.
		http.FileServer의 매개변수로 ../files/를 전달해서 ../files 디렉터리를 루트 디렉터리로 인식하도록 한다.
		따라서 ../files/save.svg 경로는 http.FileSystem에서는 /sage.svg로 인식한다.

		따라서 클라이언트가 http.FileServer로부터 sage.svg 파일을 얻어오려면 요청 경로를 /sage.svg로 해야 된다.
		하지만 각 테스트 케이스의 URL 경로에는 /static/ 접두사가 파일명 앞에 붙는다. 즉, 각 테스트는 존재하지 않는
		static/sage.svg 파일을 요청한다.
		이런 문제를 해결하기 위해서 net/http 패키지의 다른 미들웨어를 사용한다. http.StripPrefix 미들웨어는 사용자가 요청한 URL로부터
		미들웨어에 매개변수로 전달한 접두사를 제거한 후에 http.Handler로 전달한다.

		RestrictPrefix 미들웨어로 전달된 URL을 확인해보면 /static/이 제거된 형태임을 알 수 있다.

		ResticrtPrefix 미들웨어에서 문제가 발견되지 않은 경로는 http.FileServer로 요청이 전달되고
		해당하는 파일을 찾아서 응답 보디로 해당 파일의 내용을 쓰게된다.
	*/
	handler := http.StripPrefix("/static/",
		RestrictPrefix(".", http.FileServer(http.Dir("../files/"))),
	)

	testCases := []struct {
		path string
		code int
	}{
		{"http://test/static/sage.svg", http.StatusOK},
		{"http://test/static/.secret", http.StatusNotFound},
		{"http://test/static/.dir/secret", http.StatusNotFound},
	}

	for i, c := range testCases {
		r := httptest.NewRequest(http.MethodGet, c.path, nil)
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, r)

		actual := w.Result().StatusCode
		if c.code != actual {
			t.Errorf("%d:expected %d; actual %d", i, c.code, actual)
		}
	}
}
