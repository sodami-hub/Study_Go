package restrictprefix

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	"go.uber.org/zap"
)

func init() {
	caddy.RegisterModule(RestrictPrefix{})
}

// RestrictPrefix는 URI의 일부가 주어진 접두사와 일치하는 요청을 제한하는 미들웨어이다. 예를들어 요청하는 파일명 앞에 (.)이 있는 파일을 요청한다면 막는다.
type RestrictPrefix struct {
	Prefix string `json:"prefix,omitempty"`
	logger *zap.Logger
}

// CaddyModule은 Caddy의 모듈 정보를 반환한다.
func (RestrictPrefix) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.restrict_prefix",
		New: func() caddy.Module { return new(RestrictPrefix) },
	}
}

// Zap 로거를 RestrictPrefix로 프로비저닝
func (p *RestrictPrefix) Provision(ctx caddy.Context) error {
	p.logger = ctx.Logger(p)
	return nil
}

// 모듈 구성에서 접두사를 검증하고 필요시 기본 점두사를 "."으로 설정
func (p *RestrictPrefix) Validate() error {
	if p.Prefix == "" {
		p.Prefix = "."
	}
	return nil
}

// MiddlewareHandler 인터페이스 구현
// ServeHTTP는 caddyhttp.MiddlewareHandler 인터페이스를 구현
func (p RestrictPrefix) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	for _, part := range strings.Split(r.URL.Path, "/") {
		if strings.HasPrefix(part, p.Prefix) {
			http.Error(w, "Not Found", http.StatusNotFound)
			if p.logger != nil {
				p.logger.Debug(fmt.Sprintf("restricted prefix: %q in %s", part, r.URL.Path))
			}
			return nil
		}
	}
	return next.ServeHTTP(w, r)
}

/*
이 부분은 Go언어에서 타입(RestrictPrefix)이 특정 인터페이스를 구현하는지 확인하기위한 컴파일시 검사를 수행한다.
여기서 RestrictPrefix 타입이 caddy.Provisioner, caddy.Validator, caddy.MiddlewareHandler 인터페이스를 구현하는지 확인한다.
'_' 사용해서 실제 값은 무시한다. 하지만 RestrictPrefix 타입이 해당 인터페이스를 구현하지 않으면 컴파일 오류가 발생한다.
*/
var (
	_ caddy.Provisioner           = (*RestrictPrefix)(nil)
	_ caddy.Validator             = (*RestrictPrefix)(nil)
	_ caddyhttp.MiddlewareHandler = (*RestrictPrefix)(nil)
)
