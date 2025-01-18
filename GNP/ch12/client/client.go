package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"housework"
)

// 집안일 애플리케이션의 gRPC 클라이언트 초기 코드
var addr, caCertFn string

func init() {
	flag.StringVar(&addr, "address", "localhost:34443", "server address")
	flag.StringVar(&caCertFn, "ca-cert", "cert.pem", "CA certificate")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			`Usage: %s [flags] [add chore, ... | complete #]
		add			add comma-separated chores
		complete	complete designated chore
		
		Flags:
		`, filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
}

/*
gRPC는 클라이언트와 서버 간의 통신을 자동으로 처리합니다. 클라이언트가 gRPC 메서드를 호출하면,
gRPC 프레임워크가 이를 처리하여 서버의 해당 메서드로 요청을 전달하고, 서버의 응답을 클라이언트로 반환합니다.

gRPC 통신 과정
클라이언트 측:
- 클라이언트는 gRPC 클라이언트 객체를 생성하고, 이를 통해 서버의 메서드를 호출합니다.
- 클라이언트는 gRPC 프레임워크를 통해 서버와의 연결을 설정하고, 요청을 서버로 전송합니다.
서버 측:
- 서버는 gRPC 서버 객체를 생성하고, 서비스 구현체를 등록합니다.
- 서버는 클라이언트로부터 요청을 수신하고, 해당 요청을 처리할 메서드를 호출합니다.
- 서버는 처리 결과를 응답으로 클라이언트에 반환합니다.
*/

// gRPC 클라이언트를 이용하여 현재 집안일 목록 확인
func list(ctx context.Context, client housework.RobotMaidClient) error {
	chores, err := client.List(ctx, new(housework.Empty))
	if err != nil {
		return err
	}

	if len(chores.Chores) == 0 {
		fmt.Println("U have nothing to do!")
		return nil
	}

	fmt.Println("#\t[X]\tDescription")
	for i, chore := range chores.Chores {
		c := " "
		if chore.Complete {
			c = "X"
		}
		fmt.Printf("%d\t[%s]\t%s\n", i+1, c, chore.Description)
	}
	return nil
}

// gRPC 클라이언트를 이용하여 집안일 추가하기
func add(ctx context.Context, client housework.RobotMaidClient, s string) error {
	chores := new(housework.Chores)

	for _, chore := range strings.Split(s, ",") {
		if desc := strings.TrimSpace(chore); desc != "" {
			chores.Chores = append(chores.Chores, &housework.Chore{
				Description: desc,
			})
		}
	}
	var err error
	if len(chores.Chores) > 0 {
		_, err = client.Add(ctx, chores)
	}

	return err
}

// gRPC 클라이언트를 사용하여 완료된 집안일 마킹하기
func complete(ctx context.Context, client housework.RobotMaidClient, s string) error {
	i, err := strconv.Atoi(s)
	if err != nil {
		// housework.proto 코드의 chore_number 와 같이 스네이크 케이스로된 필드를
		// protoc-gen-go 명령으로 Go코드로 컴파일하면 파스칼 케이스로 바뀐다.
		// 그리고 int 타입으로 바뀐 i 값을 다시 int32값으로 바꿔서 보낸다.
		_, err = client.Complete(ctx, &housework.CompleteRequest{ChoreNumber: int32(i)})
	}

	return err
}

func main() {
	flag.Parse()

}
