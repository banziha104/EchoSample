# Echo 프레임워크

- 의존성 주입 : dep ensure -add github.com/labstack/echo@^3.1
- 기본 파일 설정 

```go
package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New() // 에코 생성
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Logger.Fatal(e.Start(":8083")) // 8083 번 포트로 전송
	
}
```

- ### [MySQL 연동](https://github.com/banziha104/EchoSample/blob/master/markdown/Mysql.md)

- ### [MongoDB 연동](https://github.com/banziha104/EchoSample/blob/master/markdown/Mongodb.md)

- ### [라우터](https://github.com/banziha104/EchoSample/blob/master/markdown/Mongodb.md)

- ### [Golang ORM](https://github.com/banziha104/EchoSample/blob/master/markdown/Gorm.md)

- ### [Request](https://github.com/banziha104/EchoSample/blob/master/markdown/Request.md)

- ### [Response](https://github.com/banziha104/EchoSample/blob/master/markdown/Response.md)