# TCP 데이터 전송하기

### net.Conn 인터페이스 사용하기
- 대부분의 네트워크 코드는 Go의 net.Conn 인터페이스 객체를 사용한다. 왜냐하면 대부분의 필요한 기능들을 net.Conn에서 제공하기 때문이다. 하위의 어떤 타입이든지 어설션할 필요없이 net.Conn 인터페이스만 이용하더라도 여러 플랫폼에서 호환할 수 있는 강인한 테스트 코드를 작성할 수 있다.
- 챕터 후반부에서는 net.Conn의 하위 타입을 어설션하여 TCP에 특정적인 고급 메서드를 사용하는 방법을 알아보겠다.
- Read와 Write 메서드 또한 net.Conn에서 유용하게 사용된다. 이 메서드는 각각 Go 언어의 표준 라이브러리와 생태계에서 범용적으로 사용되는 io.Reader 인터페이스와 io.Writer 인터페이스를 구현한다. 그 결과 io 인터페이스를 사용하는 코드들을 활용하여 강력한 기능의 네트워크 프로그램을 작성할 수 있다.
- net.Conn 인터페이스의 Close 메서드를 사용하면 네트워크 연결을 종료할 수 있다. time.Time 객체를 매개변수로 받는 setReadDeadline 메서드와 SetWriteDeadline 메서드는 매개변수로 입력받은 시간을 데드라인으로 설정하며, setDeadline 메서드는 읽기 쓰기 동시에 대해 매개변수로 입력받은 시간을 데드라인으로 결정한다. 

### 데이터 송수신
- 네트워크 연결로부터 데이터를 읽고 쓰는 것은 파일 객체에 데이터를 읽고 쓰는 것과 다르지 않다. net.Conn 인터페이스가 파일 객체에 데이터를 읽고 쓰는 데 구현된 io.ReadWriteCloser 인터페이스를 구현했기 때문이다. 
- 이번 챕터에서는 고정된 버퍼에 데이터를 읽는 방법을 알아보겠다. 그리고 네트워크 연결로부터 bufio.Scanner 메서드를 이용하여 특정 구분자를 만날 때까지 데이터를 읽는 방법을 알아보겠다. 
- 다변하는 페이로드 크기로부터 동적으로 버퍼를 할당하는 기본 프로토콜을 정의할 수 있도록 해주는 인코딩 메서드인 TLV에 대해 알아보겠다.
- 네트워크 연결로부터 데이터를 읽고 쓸 때 발생하는 에러를 처리하는 방법을 알아보겠다.

#### 데이터를 읽고 쓰는 도중 에러 처리
```
var (
    err error
    n int
    i = 7
)

for ; i>0 ; i-- {   // 일시적인 에러 발생시 재시도 하기 위한 방법
    n, err := conn.Write([]byte("hello world"))
    if err !=nil {
        nErr,ok := err.(net.Error); ok && nErr.Temporary() { // 에러가 net.Error 인터페이스를 구현했는지(ok) 그리고 에러가 일시적인지 확인(Temporary), true이면 다시 시도
            log.Println("temporary error:",nErr)
            time.Sleep(10*tiem.Second)
            continue
        }
        return err
    }
    break
}

if i==0 {
    return errors.New("temporary write failure threshold exceeded")
}

log.Printf("wrote %d bytes to %s\n",n, conn.RemoteAddr())