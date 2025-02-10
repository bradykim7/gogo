go import 정리


1. 기본 import 구문

```go
import (
    "fmt"        // 표준 라이브러리
    "net/http"   // 하위 패키지
    
    "github.com/username/repo"  // 외부 패키지
)
```

2. Import 별칭 (Alias) 사용

```go
import (
    "fmt"
    myfmt "github.com/other/fmt"    // 별칭 사용
    . "math"                        // 모든 요소를 현재 네임스페이스로 가져오기
    _ "github.com/lib/pq"          // 초기화만 하고 사용하지 않을 때
)
```

3. 모듈 시스템 (go.mod)

```go
// go.mod 파일
module github.com/username/project

go 1.21
    
require (
    github.com/pkg/errors v0.9.1
    golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
)
```

4. 의존성 관리 명령어

```shell
# 모듈 초기화
go mod init github.com/username/project

# 의존성 다운로드
go mod download

# 사용하지 않는 의존성 제거 및 필요한 의존성 추가
go mod tidy

# 의존성 버전 업데이트
go get -u github.com/pkg/errors
```

5. Import 패턴

```go
// 프로젝트 구조
myproject/
├── main.go
├── go.mod
└── internal/
└── pkg/
└── service/
└── service.go

// main.go에서 내부 패키지 import
import (
    "myproject/internal/pkg/service"
)
```

6. 특별한 Import 규칙:

```go
// 초기화만 필요한 경우 (드라이버 등)
import _ "github.com/lib/pq"

// 패키지명 충돌 방지
import (
    "database/sql"
    sqlite "github.com/mattn/go-sqlite3"
)

// 테스트 파일의 import
import "testing"  // *_test.go 파일에서 필수
```

7. Vendor 디렉토리 사용

```shell
# vendor 디렉토리 생성
go mod vendor

# vendor 디렉토리 구조
myproject/
├── vendor/
│   ├── github.com/pkg/errors/
│   └── golang.org/x/sync/
├── go.mod
└── go.sum
```


8. Import 순서 관례:

```go 
import (
// 1. 표준 라이브러리
"fmt"
"os"

    // 2. 외부 패키지
    "github.com/pkg/errors"
    "golang.org/x/sync/errgroup"
    
    // 3. 내부 패키지
    "myproject/internal/configs"
    "myproject/pkg/utils"
)
```

9. 순환 의존성 방지:

```go
// 피해야 할 패턴
// package a
package a
import "myproject/b"

// package b
package b
import "myproject/a"  // 순환 의존성! 컴파일 에러
```

실제 사용 예시:

```go
package main

import (
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    
    "myproject/internal/app"
    "myproject/internal/config"
)

func main() {
    if err := app.Run(); err != nil {
        log.Fatal(err)
    }
}
```

주요 팁:

- 항상 go mod tidy를 사용하여 의존성을 정리
- 가능한 한 적은 수의 외부 의존성 사용
- 내부 패키지는 internal 디렉토리 사용
- 표준 라이브러리를 최대한 활용
- import 경로는 항상 모듈 루트부터 시작