package json

import (
	"encoding/json"
	"io"

	"housework"
)

func Load(r io.Reader) ([]*housework.Chore, error) {
	var chores []*housework.Chore
	// Load() 메서드는 io.Reader를 매개변수로 디코더를 반환한다. 그리고 집안일 슬라이스의 포인터를 매개변수로 전달하여 디코더의 Decode 메서드를 호출한다.
	// 디코더는 io.Reader로부터 JSON 데이터를 읽어서 역직렬화한 뒤 집안일 슬라이스를 만들어낸다.
	return chores, json.NewDecoder(r).Decode(&chores)
}

func Flush(w io.Writer, chores []*housework.Chore) error {
	// Flush 메서드는 io.Writer와 집안일 슬라이스를 매개변수로 받는다.
	// 인코더의 Encode 함수는 매개변수로 집안일 슬라이스를 전달하여 JSON데이터로 직렬화한 뒤 해당 데이터를 io.Writer에 쓴다.
	return json.NewEncoder(w).Encode(chores)
}
