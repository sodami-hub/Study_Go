package main

import (
	// Caddy의 커맨드 모듈을 임포트한다. 이 빌드에는 Caddy 서버를 시작하는 main함수가 포함된다.
	cmd "github.com/caddyserver/caddy/v2/cmd"
	// Caddy의 바이너리 배포판에 존재하는 standard 모듈을 임포트
	_ "github.com/caddyserver/caddy/v2/modules/standard"

	// Caddy에 커스텀 모듈 주입하기
	_ "github.com/sodami-hub/Study_Go/GNP/ch10/caddy-restrict-prefix"
	_ "github.com/sodami-hub/Study_Go/GNP/ch10/caddy_toml_adapter"
)

func main() {
	cmd.Main()
}
