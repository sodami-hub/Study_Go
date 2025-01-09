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

// 빌드하면 caddy라는 이름의 바이너리 파일이 생성된다. 그리고 아래 명령을 실행하면 작성한 커스텀 모듈 및 어댑터가 존재함을 알 수 있다.(리눅스 및 맥OS)
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
- Caddy는 환경구성의 기본 포맷으로 JSON을 사용한다. 환경구성 정보를 JSON으로 작성해도 되지만 이미 TOML을 사용할 수 있도록 잘 동작하는 완벽한 환경구성 어댑터를 구현하였으니 이를 응용해보겠다.
- Caddy가 요청의 일부를 백엔드 웹 서비스로 리버스 프락시하도록하고, 요청의 일부는 files 서브디렉터리에서 파일을 서빙하도록 한다. 즉, 백엔드 웹 서비스의 라우트와 정적 파일의 라우트, 총 두 개의 라우트가 필요하다. 먼저 caddy.toml 파일에 서버의 환경구성 정보를 정의한다.[/caddy/caddy.toml]
```
[apps.http.servers.test_server]
listen = [
    'localhost:2020',
]
```
- TOML 어댑터는 TOML 정보를 JSON으로 그대로 변환해 줍니다. 따라서 Caddy가 기대하는 네임스페이스를 일치하게 사용할 수 있도록 주의해야 된다. 여기서 사용할 서버의 네임스페이스는 apps.http.servers.test_server이다. 이 서버는 localhost의 2020번 포트에서 연결 요청을 리스닝한다.

### 서비스로 리버스 프락시 구성하기
- Caddy는 수신 연결을 그대로 백엔드 웹 서비스로 전송할 수 있는 강력한 리버스 프락시 핸들러를 내장하고 있다. Caddy는 클라이언트의 요청과 일치하는 라우트를 찾은 후에 해당하는 라우트의 핸들러로 요청을 전달한다.
- [/caddy/caddy.toml] 파일에 라우트와 리버스 프락시 핸들러를 추가한다.

### 정적 파일 서빙하기
- 9장에서는 http.FileServer를 사용하여 정적파일을 서빙했다. Caddy에는 비슷한 기능을 하는 file_server 핸들러가 존재한다. [caddy.toml]파일에 정적 파일을 서빙하기 위한 두 번째 라우트틀 추가하겠다.


### 구성 확인해 보기
- Caddy를 시작하고 환경구성 작업이 예상대로 되었는지 확인해보겠다. 일부 이미지 파일이 존재하기 때문에 웹 브라우저의 사용을 권장한다.


```
- caddy.toml 파일과 toml 어댑터 설정을 사용하여 Caddy를 시작한다.
$ ./caddy start --config caddy.toml --adapter toml

- 이어서 백엔드 웹 서비스를 시작한다.
$ go run backend.go
```

### 자동 HTTPS 기능 추가하기
- 이제 Caddy의 핵심 기능인 자동 HTTPS 기능을 추가해보겠다. Caddy를 사용하면 잠깐 사이에 현대의 모든 웹 브라우저가 신뢰하는 인증서를 사용하여 HTTPS를 지원하는 웹 사이트를 구축할 수 있다. Caddy 서버는 굉장히 안정적으로 동작하며, 서버 동작에 전혀 영향을 주지 않고 Let's Encrypt의 인증서를 몇 달마다 주기적으로 교체해 준다. 직접 Go 기반의 웹 서버를 작성하여 이러한 동작을 구현할 수도 있지만, 검증된 웹 서버의 기능은 Caddy를 사용하는 것이 더욱 효율적이다. 개발자는 서비스에 대한 로직을 구현하는데 집중하면 되고, Caddy에 원하는 기능이 없다면 모듈을 사용하여 추가하면 된다.
- Caddy는 서빙하기 위해 설정된 도메인 네임을 확인할 수 있는 경우 자동으로 TLS를 활성화시켜준다. 이번 장에서 생성한 caddy.toml 환경 파일에는 도메인에 대한 정보가 없기 때문에 HTTPS가 활성화되지 않았다. Caddy의 소켓을 localhost로 바인딩하긴 했지만, 그건 어느 도메인에서 서빙할지를 설정한 것이 아니라 단지 소켓 주소를 리스닝하도록 한 것이다.
- 일반적으로 HTTPS를 활성화하는 방법은 Caddy의 라우트에 호스트 매처를 추가하는 것이다.
```
[[apps.http.servers.server.routes.match]]
host = [
    'example.com'
]
```
- 호스트 매처를 통해 Caddy는 현재 example.com의 도메인에서 서빙 중임을 확인할 수 있따. HTTPS를 활성화하기 위해 example.com 도메인에 대해 존재하는 TLS 인증서가 없다면 Let's Encrypt에서 도메인을 검증하고 인증서를 발급바든ㄴ다. Caddy는 필요에 따라 인증서를 관리해주고 자동으로 갱신해준다.