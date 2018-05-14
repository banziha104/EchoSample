# Go ORM(gorm)

- 설치 go get -u github.com/jinzhu/gorm

- import

```go
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
```

- 스키마 정의

```go
package schema

type ContentTable struct{
	Id int `gorm:"primary_key"`
	Content string
}

```

- db 접속부 설정

```go
package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func OrmConn() *gorm.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "echoTest"
	db, err := gorm.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		panic(err.Error())
	}
	return db
}
```

- CRUD 

```go

```