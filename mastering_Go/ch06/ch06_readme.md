# 06 - 유닉스 시스템의 작업
- Go에서의 시스템 프로그래밍 : 파일과 디렉터리 다루기, 프로세스 제어, 시그널 처리, 네트워크 프로그래밍, 시스템 및 설정 파일 입출력

### 1. stdin, stdout, stderr
- 유닉스는 프린터나 마우스를 비롯한 모든 것을 파일로 표현하기 때문에 유닉스에서 이 파일에 접근할 때는 양의 정수 값으로 표현하는 파일 디스크립터를 이용한다. 파일 디스크립터는 열린 파일에 접근하기 위한 내부 표현 방식으로, 경로를 길게 적는 것보다 훨씬 깔끔하고 간편하다. 
- 모든 유닉스 윤영체제에는 항상 열려 있는 세 가지 파일이 있다. 기본적으로 /dev/stdin, /dev/stdout, /dev/stderr 라는 세가지 특수 표준 파일명을 사용한다. 각각에 대한 파일 디스크립터는 0,1,2다. 이 세 가지 파일 디스크립터는 표준 입력, 표준 출력, 표준 에러라고 부른다.
- Go 에서 표준 입력은 os.Stdin, 표준 출력은 os.Stdout, 표준 에러는 os.Stderr로 접근한다.

### 2. 유닉스 프로세스
- 프로세스란 명령, 사용자 데이터, 시스템 데이터, 실행 시간 동안 얻은 다양한 리소스로 구성된 실행 환경이다.
- 프로그램은 명령들과 프로세스에서 초기화할 데이터로 구성된 바이너리 파일이다. 
- 유닉스 프로세스는 실행할 때마다 프로세스 ID라고 부르는 부호 없는 정수 값으로 된 고유 ID가 할당된다.
- 프로세스는 크게 사용자 프로세스, 데몬 프로세스, 커널 프로세스로 나눌 수 있다. 사용자 프로세스는 사용자 공간에서 실행하며 대개 특별한 접근 권한이 필요 없다. 데몬 프로세스는 사용자 공간에서 실행되지만 터미널과 상호작용하지 않고 백그라운드에서 구동한다. 커널 프로세스는 커널 공간에서만 실행하며, 모든 커널 데이터 구조체에 대한 접근 권한이 있다.
- Go 언어 에서는 exec 패키지를 이용해서 새로운 프로세스를 생성할 수 있지만 스레드를 관리하는 기능을 제공하지 않는다. 대신 고루틴을 사용하고 Go 런타임이 스레드 위에서 고루틴들을 관리한다.

### 3. 유닉스 시그널 처리
- 유닉스 시그널을 사용하면 애플리케이션들끼리 비동기적으로 상호작용할 수 있다. 그러나 유닉스 시그널을 처리할 때는 채널을 사용해야 한다. 시그널 처리에는 고루틴과 채널이 필요하므로 Go의 동시성 모델을 살펴보겠다.
- [/signals] 프로그램은 SIGINT와 SIGINFO를 처리한다. Go에서 해당 SIGINT는 syscall.SIGINT로 표현한다. 또한 switch 블록에서 default를 사용해 나머지 시그널들을 처리한다. switch를 통해 여러 가지 시그널에 따라 다른 동작을 수행할 수 있다. 모든 시그널을 전달받는 채널은 signal.Notify() 함수 안에 있다. 채널은 용량을 가질 수 있다. 프로그램을 종료할 때는 동시에 다른 시그널을 처리할 필요가 없으므로 용량이 1이어도 충분하다. 보통은 익명 함수를 고루틴으로 실행해 시그널을 처리한다. 시그널이 들어오면 채널로 시그널이 보내지고 다른 고루틴에서 시그널을 받은 다음 변수로 저장한다.(이때부터 채널에 다른 시그널이 들어올 수 있다.) 시그널을 저장한 변수는 switch 구문으로 처리된다.

### 4. 파일 입출력
- io.Reader, io.Writer 인터페이스, 버퍼를 이용하는 파일 입출력, bufio 패키지를 통해서 Go에서 파일을 입출력하는 방법을 살펴본다.
##### io.Reader, ioWriter 인터페이스
- io.Reader 인터페이스의 정의
```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```
- Reader 인터페이스를 만족하려면 하나의 메서드를 구현해야 한다. Read()의 매개변수는 바이트 슬라이스다. Read()의 반환값은 정수와 에러다.
- Read() 메서드는 바이트 슬라이스를 입력받아 입력된 슬라이스의 길이까지의 데이터를 채운 뒤 읽은 바이트의 개수와 에러를 함께 반환하는 함수다.
- io.Writer 인터페이스의 정의
```
type Writer interface {
    Write(p []byte)(n int, err error)
}
```
- Write() 메서드는 파일에 쓰고자 하는 바이트 슬라이스를 받아서 쓴 바이트의 개수와 에러를 반환한다.

##### io.Reader와 io.Writer의 사용과 오용 [/ioInterface]
##### 버퍼를 이용한 파일 입출력과 버퍼를 이용하지 않는 파일 입출력
- 버퍼를 이용한 파일 입력과 출력은 데이터를 읽거나 쓰기 전에 잠시 버퍼에 저장한다. 따라서 파일은 한 바이트 단위로 읽지 않고 한 번에 여러 바이트를 읽을 수 있다. 데이터를 버퍼에 저장해둔 후 각자 원하는 방식으로 읽는 것이다.
- 버퍼를 사용하지 않는 입출력은 실제로 파일을 읽거나 쓰기 전에 데이터를 임시로 저장하지 않는다.
- 중요한 데이터를 다룰 때는 일반적으로 버퍼를 사용하지 않는 것이 좋다. 버퍼를 거치는 동안 데이터가 더 이상 쓸모없는 상태가 되거나 버퍼에 저장된 사이에 컴퓨터 전원이 꺼지면 데이터를 잃어버릴 수 있기 때문이다. 하지만 버퍼 사용에 대한 기준은 명확하게 하기 어렵다. 따라서 구현하기 쉬운 방식을 사용해야 된다. 
- 그러나 버퍼를 사용하면 파일이나 소켓으로부터 데이터를 읽을 때 시스템 콜을 호출하는 횟수가 줄어들어 성능이 나아진다는 점은 알고 있어야 된다. 
- 또한 bufio 패키지도 있다. bufio 패키지는 버퍼를 사용한 입력과 출력 기능을 제공한다. bufio 패키지는 텍스트 파일을 읽는 데에 매우 널리 사용하며 다음 절에서 알아보겠다.

### 5. 텍스트 파일 읽기
- 텍스트 파일을 읽는 방법과 난수를 얻는 데 /dev/random 유닉스 장치를 사용하는 방법을 알아본다.
##### 줄 단위로 텍스트 파일 읽기[/lineByLine]
##### 단어 단위로 텍스트 파일 읽기[/wordByWord]
- 정규표현식을 사용해서 각 줄에 있는 단어를 구분한다. 전규표현식은 regexp.MustCompile("[^\\s]+")로 정의한다. 공백 문자를 구분자로 사용해 단어를 나눈다.
##### 문자 단위로 텍스트 파일 읽기[/charByChar]
- 읽어 들인 각 줄을 for 루프와 range를 이용해서 반복한다. 
##### /dev/random 읽기[/devRandom]
- /dev/random 시스템 장치는 랜덤 데이터를 생성하고자 존재하며 이로부터 만들어진 결과는 프로그램을 테스트하거나 난수 생성기의 시드로 활용할 수 있다.
- /dev/random 에서 데이터를 가져오는 것은 약간 까다롭기에 여기서 따로 설명한다.
##### 파일에서 원하는 만큼만 데이터 읽기[/readSize]
- 파일의 작은 일부분만을 보고 싶은 경우에 사용할 수 있는 코드이다.

### 6. 파일 쓰기
- [/writeFile]

### 7. JSON 데이터 처리
##### Marshal(), Unmarshal()[/encodeDecode]
- Go에는 JSON데이터를 다루기 위한 표준 라이브러리인 encoding/json 라이브러리가 있다. 또한 Go는 태그를 이용해 구조체에서 JSON 필드를 지원할 수 있다.(이후에 살펴본다.) 태그를 이용하면 JSON 레코드를 인코딩하거나 디코딩 할 수 있다. 그러나 먼저 JSON 레코드를 마샬링하거나 언마샬링하는 방법부터 알아보겠다.
- 마샬링(Marshaling) : 구조체를 JSON 레코드로 변환하는 과정이다. 보통 JSON레코드 형태로 네트워크를 통해 데이터를 전송하거나 디스크에 저장할 때 사용한다.
- 언마샬링(Unmarshaling) : 바이트 슬라이스 형태의 JSON 레코드를 구조체로 변환하는 과정이다. 네트워크를 통해 받거나 디스크에 파일로 저장한 JSON 데이터를 코드에서 사용할 때 언마샬링이 필요하다.
##### 구조체와 JSON[/tagsJSON]
- 빈 필드는 JSON에서 제외하고 싶다면 다음 코드처럼 omitempty를 사용한다.
```
type NoEmpty struct {
    Name string `json:"username"`
    Year int    `json:"creationyear,omitempty"`
}
```
- 민감한 데이터를 갖는 필드는 JSON 레코드에 포함시키고 싶지 않다면 "-"값을 사용한다.
```
type NoEmpty struct {
    Name string `json:"username"`
    Year int    `json:"creationyear,omitempty"`
    password string `json:"-"`
}
```
##### 스트림 형태로 JSON 데이터 읽고 쓰기[/JSONstreams]
- Go는 여러 JSON 레코들르 처리할 수 있도록 스트림 형태로 JSON 데이터를 처리할 수 있는 방법을 제공하는데, 더 빠르고 효율적이다. 
##### JSON 레코드 출력 다듬기[/prettyPrint]
- json의 인코딩 또는 마샬에서 indent를 사용하는 방법을 보여준다.

### 8. XML 데이터 처리[/xml]
- Go 구조체에 XML 관련 태그를 추가해 XML 레코드를 encoding/xml 패키지의 xml.Unmarshal()과 xml.Marshal()을 이용해 관련 필드를 직렬화하거나 역직렬화한다. 하지만 /xml/xml.go 에서 JSON과의 차이점을 살펴볼 수 있다.
##### JSON과 XML변환[/JSON2XML]

### 9. YAML 데이터 처리[/yaml]
- Go 표준 라이브러리는 YAML 파일을 지원하지 않으므로 YAML을 지원하는 외부 라이브러리가 있는지 살펴봐야 된다. 다음 세개의 패키지들이 Go에서 YAML을 다룰 수 있게 해준다.
- https://github.com/kylelemons/go-gypsy
- https://github.com/go-yaml/yaml
- https://github.com/goccy/go-yaml

### 10. viper 패키지

### 11. cobra 패키지

### 12. 유닉스 파일 시스템에서 순환 참조 찾기