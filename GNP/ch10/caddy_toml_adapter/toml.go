package tomladapter

import (
	"encoding/json"

	"github.com/caddyserver/caddy/v2/caddyconfig"
	"github.com/pelletier/go-toml"
)

func init() {
	// 환경구성 어댑터를 Caddy에 등록한다. 이를 위해서 어댑터의 타입("toml")과 caddyconfig.Adapter 인터페이스를 구현하는 Adapter 객체를
	// caddyconfig.RegisterAdapter 함수에 매개변수로 넣어서 init 함수에서 호출한다.
	caddyconfig.RegisterAdapter("toml", Adapter{})
}

// 어댑터는 TOML 형식의 Caddy 구성 파일을 JSON으로 변환
type Adapter struct{}

// TOML형식의 보디를 JSON으로 변환
func (a Adapter) Adapt(body []byte, _ map[string]interface{}) ([]byte, []caddyconfig.Warning, error) {
	tree, err := toml.LoadBytes(body) // 토마스 펠레티어의 go-toml 라이브러리를 사용하여 환경구성 파일의 내용을 파싱한다.
	if err != nil {
		return nil, nil, err
	}

	b, err := json.Marshal(tree.ToMap()) // 파싱된 TOML을 맵으로 변환하고 맵을 JSON으로 마샬링한다.

	return b, nil, err
}
