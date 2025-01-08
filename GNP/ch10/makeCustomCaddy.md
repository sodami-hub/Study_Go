# Caddy로 웹 서버 만들어보기
## 모듈과 어댑터로 Caddy확장하기
Caddy는 기능들을 확장하기 위하여 모듈형 구조를 채택하였다. 모듈형 구조로 인하여 모듈과 환경구성 어댑터를 작성하여 직접 Caddy의 기능을 확장할 수 있다. 이번 섹션에서는 환경 구성 어댑터를 직접 작성하여 Caddy의 환경 구성 정보를 TOML(Tom's Obvious, Minimal Language) 파일에 저장한다. 또한 이전 장의 restrict_prefix 미들웨어를 모듈로 만들어보겠다.

### 환경구성 어댑터 작성하기
- Caddy는 환경구성을 JSON 포맷에서 TOML과 같은 다른 포맷으로 적용할 수 있는 어댑터 기능을 지원한다. TOML에는 주석과 멀티라인 문자열을 지원한다. Caddy는 환경구성 정보를 저장하기 위한 Caddyfile이라는 이름의 
커스텀 포맷을 지원한다. 파일 명이 Caddyfile인 경우 Caddy는 caddyfile 어댑터를 자동으로 사용한다. 커맨드 라인에서 직접 어댑터를 지정해 주려면 다음과 같이 사용한다.
```
$ caddy start --config Caddyfile --adapter caddyfile
```
- adapter 플래그는 Caddy에게 어떤 어댑터를 사용해야 하는지 명시한다. Caddy는 어댑터를 실행하여 입력받은 환경구성 정보를 JSON으로 변환하고, 어댑터가 반환한 JSON 데이터를 파싱해서 사용한다.
- [/caddy_toml_adapter] 위치에 TOML 환경구성 어댑터를 작성해보겠다.

## 접두사를 제한하는 미들웨어 모듈 작성하기
앞선 챕터에서 서버가 요청을 받았을 때 요청과 응답을 조작하고 요청의 세부 사항을 로깅하는 등의 부가적인 작업을 처리할 수 있도록 해주는 디자인 패턴인 미들웨어의 개념에 대해 살펴보았다. 
Caddy에서 미들웨어를 사용하는 방법을 알아보겠다
Go에서 미들웨어는 http.Handler를 매개변수로 받아서 http.Handler를 반환하는 함수이다.
```
func(http.Handler) http.Handler
```
어떤 객체가 ServeHTTP 메서드를 구현하면 그 객체는 http.Handler 인터페이스의 구현체이다. ServeHTTP 메서드는 http.ResponseWriter와 http.Request를 매개변수로 받는다.
```
type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}
```
Go 에서는 사실상 미들웨어와 핸들러와의 구조적인 차이는 존재하지 않는다. 개발자가 함수를 어디에 쓰냐에 따른 역할의 차이라고 할 수 있겠다.

하지만 Caddy의 미들웨어는 이러한 패턴으로 사용할 수 없다. 따라서 RestrictPrefix 함수를 그대로 사용할 수 없다. Caddy에는 핸들러에 대한 인터페이스만 존재하는 net/http와는 달리 
핸들러와 미들웨어에 대한 인터페이스가 둘 다 존재한다. http.Handler 인터페이스의 Caddy 버전으로 caddyhttp.Handler가 있다.
```
type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request) error
}
```
위에서 보는 바와 같이 caddyhttp.Handler 인터페이스의 ServeHTTP 메서드는 error 인터페이스를 반환하는 차이가 있다.

Caddy의 미들웨어는 caddyhttp.MiddlewareHandler 인터페이스를 구현하는 특별한 종류의 핸들러이다.
```
type MiddlewareHandler interface {
    ServeHTTP(http.ResponseWriter, *http.Request, Handler) error
}
```
Caddy는 미들웨어가 요청과 응답을 매개변수로 받고, 미들웨어 이후에 처리될 핸들러도 같이 매개변수로 받는다.

### Caddy에 모듈 주입하기
- 작성한 모듈과 어댑터 모두 초기화 시에 스스로 등록한다. 이 기능을 사용하기 위해 해야 할 일은 Caddy를 빌드할 때 임포트하는 것이다. 이를 위해 소스코드로부터 Caddy를 컴파일해야 한다. 먼저, 빋드를 위한 디렉터리를 생성한다.
```
$ mkdir caddy
$ cd caddy
```
- [/caddy/main.go]소스코드에서 Caddy를 빌드하기 위해서는 모듈 내에 포함시킬 소량의 보일러플레이트 코드를 필요로 한다. 모듈은 임포트의 결과로 스스로 등록하게 된다. main.go 로 새로운 파일을 만들고 코드를 작성한다. 코드에대한 설명은 코드에 있다.
- 그리고 caddy 모듈을 초기화하고 빌드한다.
```
$ go mod init caddy
$ go mod tidy
$ go build main.go

// 빌드하면 caddy라는 이름의 바이너리 파일이 생성된다. 그리고 아래 명령을 실행하면 작성한 커스텀 모듈 및 어댑터가 존재함을 알 수 있따.
$ ./caddy list-modules | grep "toml\|restrict_prefix"
caddy.adapters.toml
http.handlers.restrict_prefix
```
- caddy 바이너리를 사용하면 TOML 파일에서 환경구성 정보를 읽을 수 있고, 주어진 접두사를 가진 클라이언트의 리소스 요청을 거절할 수 있다.

## 백엔드 웹 서비스로 요청 리버스 프락시하기
- 이제 Caddy에서 무언가 의미 있는 것을 만들기 위한 기본 준비는 다 되었다. Caddy를 사용해서 여태까지 배운 걸 통합 구성하여 클라이언트의 요청을 백엔드 웹 서비스로 리버스 프락시하고, 정적 파일을 서빙해 보겠다. 두 개의 엔드포인트를 만든다. 
1. 첫 번째 엔드포인트는 Caddy의 파일 서버에서 정적 파일을 서빙하는 엔드포인트이고 
2. 두 번째 엔드포인트는 클라이언트의 요청을 백엔드 웹 서비스로 리버스 프락시하는 엔드포인트이다. 이 백엔드 웹 서비스는 클라이언트에게 HTML을 전송하며, 이 HTML을 렌더링하기 위한 리소스는 Caddy의 파일 서버에서 서빙할 것이다.
```
// 먼저 두개의 서브 디렉터리를 만들고
$ mkdir files backend
```
- files에는 정적파일들을 backend 에는 간단한 백엔드 서비스 코드를 저장한다.

### 간단한 백엔드 웹 서비스 작성하기
- Caddy가 사용자의 요청을 리버스 프락시할 백엔드 웹 서비스가 필요하다. 이 서비스는 모든 요청에 대해 HTML 문서를 응답한다. HTML 문서 내의 리소스는 백엔드 웹 서비스 대신 Caddy에서 서빙한다.
- [/backend/main.go]

### Caddy 환경구성 설정하기
- 
