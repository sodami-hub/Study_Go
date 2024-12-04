# ch08 - 웹 서비스 만들기
- net/http 패ㅔ키지를 이용해 http 프로토콜을 사용해본다. 
- 파일 전송 프로토콜 서버를 만드는 방법
- Go 애플리케이션에서 발생한 메트릭을 프로메테우스로 전달하는 방법
- runtime/metrics 패키지를 이용해 Go 런타임에서 메트릭을 가져오는 방법

### net/http 패키지
- net/http패키지는 웹 서버와 클라이언트를 개발하는 데 필요한 기능을 제공한다. 예를 들어 http.Get()과 http.NewRequest()를 이용하면 클라이언트에서 HTTP 요청을 보낼 수 있고 http.ListenAndServer() 함수를 이용하면 명시된 IP주소와 TCP 포트로 웹 서버를 시작할 수 있다. 또한 http.HandleFunc()를 통해 지원하는 URL 및 해당 URL로 들어오는 요청을 처리할 수 있는 함수를 정의할 수 있다.

##### http.Response 타입
- http.Response는 HTTP 요청에 대한 응답을 표현하는 타입이다. http.Client와 http.Transport는 응답 헤더가 도착할 때 http.Response 값을 반환한다. 
```
type Response struct {
    Status      string  // e.g. "200 OK"
    StatusCode  int     // e.g. 200
    Proto       string  // e.g. "HTTP/1.0"
    ProtoMajor  int
    ProtoMinor  int
    Header      Header
    Body        io.ReadCloser
    ...
    ...
    TLS         *tls.ConnectionState
}
```
- 위 구조체의 모든 필드에 대해서 알 필요는 없다. 하지만 Status, StatusCode, Body와 같은 필드는 중요하므로 잘 알아둬야 된다.
##### http.Request 타입
- 서버가 받거나 클라이언트가 보내는 요청을 표현하는 타입이다. http.Request 타입의 퍼블릭 필드는 다음과 같다.
```
type Request struct {
    Method      string
    URL         *url.URL
    Proto       string
    ...
    Body        io.ReadCloser
    ...
    ...
    Response    *Response
}
```
- Body 필드는 요청의 본문을 담고 있다. 요청의 본문을 읽은 뒤에는 GetBody()를 호출해 본문의 복사본을 만들 수 있다.
##### http.Transport 타입
- http.Transport 타입을 사용하면 HTTP 연결을 더 세부적으로 제어할 수 있다. 따라서 정의가 길고 복잡하다. http.Transport는 상당히 저수준의 구조체다. 반면 이번 장에서 사용할 http.Client는 고수준의 HTTP 클라이언트이고 각각의 http.Client는 Transport 필드를 갖고 있다. Transport 필드값이 nil이라면 DefaultTransport가 사용된다.(go doc http.DefaultTransport)

### 웹 서버 만들기
- Go 웹 서버에서도 많은 일을 효율적이고 안전하게 수행할 수 있다. 하지만 모듈, 다수의 웹 사이트, 가상 호스트 등을 지원하는 서버가 필요하다면 아파치, 엔진엑스나 Go로 만들어진 캐디(Caddy)와 같은 서버를 사용하는 편이 낫다.
- HTTPS가 아닌 HTTP를 사용하는 이유는 Go 웹 서버는 도커 이미지 형태로 배포돼 HTTPS로 안전한 인증을 제공하는 웹 서버 뒤에 숨겨져 있기 때문이다. 또한 애플리케이션을 어떤 도메인 이름으로 배포할지 모르는 상태에서 HTTPS 프로토콜을 사용할 수는 없다. 
- net/http 패키지는 웹 서버와 클라이언트를 개발할 수 있는 함수와 데이터 타입을 제공한다. http.Set()과 http.Get() 메서드는 HTTP 및 HTTPS 요청을 생성하는 데 사용할 수 있고 http.ListenAndServe()는 들어오는 요청을 사용자 정의 핸들러 함수를 이용해 처리할 수 있게 만들어준다. http.HandleFunc()를 사용하면 지원할 엔드포인트 및 클라이언트 요청을 처리할 핸들러 함수를 정의할 수 있다. 또한 이 함수는 여러 번 호출될 수 있다.
- [wwwServer.go]

### 전화번호부 애플리케이션 업데이트


### 프로메테우스로 메트릭 노출



### 웹 클라이언트 개발



### 전화번호부 서비스를 위한 클라이언트 만들기


### 파일 서버 만들기



### HTTP 연결 타임아웃