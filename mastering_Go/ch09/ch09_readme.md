# TCP/IP와 웹소켓
9장에서는 좀 더 저수준의 프로토콜인 TCP와 UDP 프로토콜을 살펴보고 net 패키지를 이용해 TCP/IP 서버와 클라이언트를 개발하는 방법을 알아본다. 또한 HTTP에 기반을 둔 웹소켓 프로토콜과 로컬 환경에서만 동작하는 유닉스 도메인 소켓을 이용해 서버와 클라이언트를 개발하는 방법도 알아본다.

## TCP/IP
- TCP/IP란 인터넷을 구성하는 프로토콜 집합을 말한다. 이 이름은 가장 널리 알려진 프로토콜인 TCP와 IP에서 따온 것이다. TCP란 전송 제어 프로토콜(Transmission Control Protocol)의 줄임말이다. TCP 에서는 데이터를 세그먼트란 단위로 전송하고, 이를 TCP 패킷이라고도 부른다. TCP의 큰 장점은 신뢰성을 보장한다는 것이다.
- 두 머신 사이에 TCP 연결이 생성되면 전화를 연결할 때처럼 전이중(full duplex) 방식의 가상 회로가 생성된다. 그런 다음엔 두 머신 사이에 주고받는 데이터가 정확하다는 것이 보장되는 상태에서 지속적으로 통신할 수 있다. TCP 패킷의 헤더에는 출발지 포트와 목적지 포트가 있다. 이 두개의 필드와 출발지 IP주소와 목적지 IP주소를 한데 묶어 하나의 TCP 연결에 대한 고유한 식별자로 사용한다.
- 🙌0-1024번 포트 번호는 루트(root) 사용자만 접근할 수 있다는 것을 참고하자.
- iP는 인터넷 프로토콜의 줄임말이다. IP의 가장 큰 특징은 자체적으로는 신뢰성을 보장하지 않는다는 것이다. 
- UDP는 IP를 기반으로 작동하며 신뢰성을 보장하지 않는다. 
#### nc(1) 커맨드라인 유틸리티
-  nc(1) 유틸리티는 netcat(1)이라고도 부르며 TCP/IP 서버와 클라이언트를 테스트하는 데 굉장히 유용하다. 사실 nc(1)은 TCP 연결 맺기, UDP 메시지 주고받기, TCP 서버 역할 등 TCP, UDP, IPv4, IPv6와 관련된 모든 기능을 수행할 수 있다.
- nc(1)을 TCP 서비스의 클라이언트로도 사용할 수 있다. 예를 들어 TCP 서비스의 IP주소가 10.10.1.123이고 1234번 포트를 갖고 있다면 다음과 같은 커맨드를 실행해 TCP연결을 맺을 수 있다.
```
$ nc 10.10.1.123 1234
```
- -l 옵션을 지정하면 netcat(1)을 서버로 구동할 수 있고 커맨드를 실행하면 netcat(1)에 지정한 포트 번호로 연결 요청이 들어오길 기다리게 된다. 기본적으로 nc(1)는 TCP 프로토콜을 사용한다. 그러나 -u 옵션을 사용하면 UDP 프로토콜을 사용한다. 마지막으로 -v와 -vv 옵션을 지정하면 netcat(1)의 출력을 상세하게 표현한다.

## net 패키지
- Go 표준 라이브러리인 net 패키지는 TCP/IP, UDP, 도메인 이름 확인, 유닉스 도메인 소켓을 다루는 패키지다. net.Dial() 함수는 클라이언트에서 네트워크에 연결할 때 사용하고 net.Listen()함수는 서버 역할을 하는 Go 프로그램에서 네트워크 접속을 받는데 사용한다. net.Dial() 함수와 net.Listen() 함수는 모두 net.Conn 타입의 값을 반환한다. 이 타입은 io.Reader, io.Writer인터페이슬르 구현한 타입이므로 파일 입출력과 관련된 코드를 이용해 net.Conn 연결에서 데이터를 읽거나 쓸 수 있다. 


## TCP 클라이언트
#### net.Dial()을 이용한 TCP 클라이언트 개발[/tcpC]
#### net.DialTCP()를 이용해 TCP 클라이언트 개발[/otherTCPclient]

## TCP 서버
TCP 서버를 통해 TCP 클라이언트와 상호작용하는 두 가지 방식을 살펴본다.
#### net.Listen()을 이용한 TCP 서버 개발
- net.Listen()을 이용해 형재 날짜와 시각을 네트워크 패킷 하나에 담아 클라이언트에 반환하는 TCP 서버를 개발해본다. 구체적으로 클라이언트와 연결한 후에 운영체제로부터 날짜와 시각을 알아내 그 값을 다시 클라이언트에 보낸다 net.Listen() 함수는 연결이 맺어지길 기다리는 데 사용하고 net.Accept() 메서드는 다음 연결이 들어올 때까지 기다리다가 연결이 맺어지면 클라이언트 정보와 함께 Conn 변수를 반환한다.[/tcpS]
#### net.ListenTCP()를 사용한 TCP 서버 개발
- [/otherTCPserver.go]TCP 서버는 에코 서비스를 제공하고 이는 클라이언트에서 수신한 데이터를 그대로 돌려주는 서비스다.

## UDP 클라이언트[/udpC]

## UDP 서버[/udpS]

## 동시성을 지원하는 TCP 서버[/concTCP]
- 동시성있는 TCP 서버에서는 Accept()를 호출할 때마다 새로운 고루틴을 이용해 클라이언트에 서비스를 제공한다. 따라서 하나의 서버가 동시에 여러 클라이언트에 서비스를 제공할 수 있다. 실제 프로덕션 환경의 서버들은 항상 이런 방식으로 구현한다.

## 유닉스 도메인 소켓
- 유닉스 도메인 소켓 또는 프로세스 간 통신 소켓은 같은 머신에서 실행하는 프로세스 사이에서 데이터를 교환하고자 사용하는 데이터 통신 엔드포인트다. 
같은 머신에서는 TCP/IP연결 대신 유닉스 도메인 소켓을 사용한다. 그 이유는 빠르고, 상대적으로 더 적은 자원을 사용하기 때문이다.
#### 유닉스 도메인 소켓 서버[/socketServer]
#### 유닉스 도메인 소켓 클라이언트[/socketClient]

## 웹소켓 서버 만들기
- 웹소켓 프로토콜은 하나의 TCP 연결 위에서 전이중 통신(동시에 양뱡향으로 데이터를 전송할 수 있다.) 채널을 제공하는 통신 프로토콜이다. 웹소켓 프로토콜은 RFC6455에 정의돼 있으며 http://, https:// 대신 ws://, wss:// 를 사용한다. 이번 절에서는 gorilla/websocket 모듈을 상ㅇ해 작지만 제기능을 하는 웹소켓 서버를 개발해본다. 서번느 에코 서비스를 구현해 클라이언트의 입력을 그대로 반환하게 만들 것이다.
- golang.org/x/net/websocket 패키지를 사요해도 웹소켓 클라이언트와 서버를 개발할 수 있다. 하지만 해당 패키지에는 일부 기능이 부족하므로 https://godoc.org/github.com/gorilla/websocket이나 https://godoc.org/nhooyr.io/websocket을 사용하는 것을 권장한다.
- 웹소켓 프로토콜의 장점은 다음과 같다.
1. 전이중 통신을 제공하는 양방향 채널이다. 따라서 서버는 클라이어늩가 데이터를 읽기를 기다릴 필요가 없고, 클라이언트도 마찬가지다.
2. TCP 소켓이므로 HTTP 연결을 맺기 위한 추가적인 비용이 발생하지 않는다.
3. 웹소켓 연결은 HTTP 데이터를 전송하는데 사용할 수 있지만, 일반 HTTP 연결이 웹소켓처럼 동작하지 않는다.
4. 웹소켓 연결은 끊어지지 않고 계속 유지할 수 있으므로 매번 새로 열 필요가 없다.
5. 웹소켓 연결은 실시간 웹 애플리케이션에서 사용할 수 있다.
6. 클라이언트가 요청하지 않아도 서버에서 클라이언트로 어느 때나 데이터를 전송할 수 있다.
7. 웹소켓은 HTML5 규격에 포함돼 있기 때문에 모든 현대 웹 브라우저에서 지원한다.

- 서버 구현을 살펴보기 전에 gorilla/websocket 패키지의 websocket.Upgrader메서드를 알아둘 필요가 있다. websocket.Upgrader 메서드는 HTTP 서버 연결을 웹소켓 프로토콜로 업그레이드하고 업그레이드를 위한 캐개변수를 지정할 수 있다. 그런 다음 HTTP 연결은 웹소켓 연결이 돼 HTTP 프로토콜의 메시지는 보낼 수 없게 된다.
##### 서버 구현[/ws]

## 웹소켓 클라이언트 만들기[/wsClient]


## 연습문제
1. 동시성을 지원하며 미리 정의한 특정한 범위의 난수들을 생성해주는 TCP 서버를 개발해보자
2. 동시성을 지원하며 클라이언트에서 지정한 범위의 난수들을 생성해주는 TCP 서버를 개발해보자. 값들이 저장된 배열에서 랜덤하게 값을 지정하는 방식으로도 구현할 수 있다.
3. 동시성 있는 TPC 서버에서 유닉스 시그널을 처리하는 기능을 추가해 우아한 종료가 가능하도록 개발해보자
4. 난수를 생성하는 유닉스 도메인 소켓 서버를 개발해보자. 그런 다음 해당 서버의 클라이언트도 개발해보자
5. 클라이언트가 보낸 개수만큼의 랜덤한 정수를 생성하는 웹소켓 서버를 개발해보자. 생성할 정수의 개수는 클라이언트의 첫 번째 메시지에서 주어진다.

