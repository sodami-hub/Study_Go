# Go language를 위한 Repository.
## Network, distribution, SDN... 을 향한 한걸음

### Go_Basic_Programming

### GNP => Go_Network_Programming

### Mastering_Go

##### 로컬 환경에서의 모듈 임포트 방법
- GNP/ch09/server.go 의 import 문과 go.mod 파일을 참고하라
```
// server.go 의 import 문
import (
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"GNP/handlers"  // 이 패키지를 임포트해야 된다.
)

// 해당 패키지의 go.mod 파일
module servers

go 1.23.2

replace GNP/handlers => ../handlers // 이렇게 추가해주고 go mod tidy 하면 된다. 
// 당연히 GNP/handlers 패키지에도 go.mod 파일이 있어야 된다.

```