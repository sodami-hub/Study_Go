package protobuf

import (
	"io"

	// JSON, Gob은 housework 패키지에 의존해서 구현했지만 protoc 컴파일러를 통해서
	// go로 컴파일된 패키지명인 v1을 임포트한다.
	"github.com/sodami-hub/Study_Go/GNP/ch12/housework/v1"
)

func Load(r io.Reader) ([]*housework.Chore, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var chores housework.Chores
}
