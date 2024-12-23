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

### 전화번호부 애플리케이션 업데이트 [/www-phone-main]
- go run www-phone.go handlers.go
- localhost:1234 // url
- localhost:1234/list
- localhost:1234/status
- localhost:1234/search/[전화번호]
- localhost:1234/insert/[성]/[이름]/[전화번호]
- localhost:1234/delete/[전화번호]

### 프로메테우스로 메트릭 노출
- 파일을 디스크에 쓰는 애플리케이션이 있다고 가정하고 해당 애플리케이션에서 나오는 메트릭들을 통해 여러 파일에 데이터를 쓰는 행위가 성능에 얼마나 영향을 끼치는지 알고 싶다고 가정해보자. 이 경우 애플리케이션의 행동을 더 잘 이해할 수 있게 성능 데이터를 수집할 필요가 있다. 예제에서는 게이지 타입의 메트릭만 다루지만 프로메테우스는 여러 타입의 메트릭 데이터를 수집할 수 있다. 프로메테우스에서 지원하는 메트릭의 타입은 다음과 같다.
```
- 카운터 : 증가하는 카운터를 나타내는 데 사용되는 누적 값이다. 카운터의 값은 증가, 유지, 0으로 초기화 할 수 있지만 감소시킬 수 없다. 지금까지의 요청 수, 에러 수와 같은 누적 값을 나타내는데 사용된다.
- 게이지 : 증가 또는 감소할 수 있는 단일 숫자 값이다. 게이지는 보통 요청의 수, 지속 시간 등과 같이 증가하거나 감소할 수 있는 값을 나타내는 데 사용한다.
- 히스토그램 : 관측값을 샘플링하고 카운트 및 버킷을 만드는 데 사용한다. 히스토그램은 보통 요청 지속 시간, 응답 시간과 같은 값에 사용한다.
- 요약: 히스토그램과 비슷하지만 시간에 따른 슬라이딩 윈도우에서 분위를 계산할 수도 있다.
```
- 히스트그램과 요약은 통계 계산을 하는 데 유용하다. 하지만 시스템 메트릭을 저장할 때는 보통 카운터나 게이지 타입만을 사용한다.
- 예제를 통해서 메트릭을 수집해 이를 프로메테우스에 노출하는 방법을 알아본다. 편의상 난수를 생성하는 애플리케이션을 활용하겠다. 먼저 Go 런타임과 관련된 메트릭을 제공하는 runtime/metrics 패키지를 알아보겠다.
##### runtime/metrics 패키지
- 해당 패키지는 Go 런타임의 메트릭을 개발자가 사용할 수 있게 만들어준다. 각각의 메트릭 이름은 경로처럼 정의한다. 예를 들어 /sched/goroutines:goroutines 메트릭은 고루틴의 개수를 알 수 있는 메트릭이다. 모든 메트릭을 수직하고 싶다면 metrics.All()을 사용해야 한다. 메트릭은 metrics.Sample 데이터 타입에 저장한다.
```
type Sample struct {
    Name string
    Value Value
}
```
- Name 값은 metrics.All()에서 반환하는 메트릭들의 값 중 하나여야 한다. 메트릭에 대한 이름을 이미 알고 있다면 metrics.All()을 사용할 필요가 없다. 
- [/metrics] 에서는 Go에서 제공하는 runtime/metrics 패키지를 사용해서 메트릭 값을 확인하겠다. /sched/goroutines:goroutines의 값을 가져와 화면에 출력한다.
##### 메트릭 노출[/samplePro]
- 메트릭을 수집하는 것과 이를 프로메테우스가 수집할 수 있게 하는 것은 완전히 다른 작업이다. 이 절에서는 메트릭을 프로메테우스가 수집할 수 있게 만드는 방법을 살펴본다.
##### Go 서버의 도커 이미지 생성
- 도커 이미지로 얻을 수 있는 가장 큰 장점은 컴파일러나 필요한 리소스가 존재하는지 등의 걱정 없이 도커 환경에 배포할 수 있다는 것이다. 모든 것이 도커 이미지에 포함돼 있기 때문이다.
- 도커 이미지 대신 일반 Go 바이너리를 사용하면 되지 않는가? 라는 질문을 할 수 있다. 답은 간단하다. 도커 이미지는 docker-compose.yml 파일에 들어갈 수도 있고, 쿠버네티스를 이용해 배포할 수도 있기 때문이다. Go 바이너리는 이와 같이 활용할 수 없다.


### 웹 클라이언트 개발
- [/simpleClient] http.Get(URL) 호출을 통해서 요청을 보내고 응답을 가져온다. 
- [/wwwClient] http.NewRequest()를 사용해 클라이언트 개선. 앞의 내용은 상대적으로 간단하지만 아무런 유연성도 제공하지 않는다. 따라서 이번에는 http.Get()함수를 사용하지 않고 URL을 읽는 방법과 더 많은 옵션을 알아본다. 더 많은 코드를 작성해야 된다.

### 전화번호부 서비스를 위한 클라이언트 만들기


### 파일 서버 만들기



### HTTP 연결 타임아웃
- 서버나 클라이언트에서 시간이 너무 오래 결리는 네트워크 연결을 타임아웃시키는 방법을 알아본다.
##### SetDeadline() 사용
- net 패키지에서 사용하는 SetDeadline() 함수는 주어진 네트워크 연결에서 읽기와 쓰기를 수행하는 데 적용하는 데드라인을 설정한다. SetDeadline() 함수의 작동 방식 특성상 읽기나 쓰기 연산을 수행하기 전에 SetDeadline()을 호출해야 한다. 여기서 명심할 점은 Go에서는 타임아웃을 구현할 때 데드라인 방식을 이용한다는 점이다. 따라서 애플리케이션이 데이터를 보내거나 받을 때마다 이 값을 리셋할 필요는 없다. SetDeadline()을 사용하는 방법은 [/withDeadline]에 나와있고 Timeout() 함수에 구현돼 있다.
##### 클라이언트에서 타임아웃 설정[/timeoutClient]
- 클라이언트에서 너무 긴 시간 동안 종료되지 않는 네트워크 연결을 타임아웃시키는 방법을 소개한다. 클라이언트가 정해진 시간동안 서버로부터 응답을 받지 못하면 클라이언트에서 연결을 종료한다.
##### 서버에서 타임아웃 설정[/timeoutServer]
- 너무 긴 시간 동안 종료되지 않는 네트워크 연결을 타임아웃시키는 방법을 소개한다. 클라이언트보다 서버에서의 타임아웃이 더 중요한데, 서버에 너무 많은 수의 연결이 맺어져 있을 경우 그 연결들이 닫히기 전까지 다른 요청을 처리할 수 없을 수도 있기 때문이다. 이런 현상이 발생하는 데는 크게 두가지 원인이 있다. 하나는 클라이언트 소프트웨어에 버그가 있을 때고 다른 하나는 서버 프로세스가 서비스 거부 공격을 받을 때다.