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
// Create
	ormRouter.POST("/create", func(c echo.Context) error {
		controller.CreateORMDataSnippet(schema.ContentTable{4,"content4"},mysql.OrmConn())
		return c.String(http.StatusOK,"인설또")
	})

	// Read
	ormRouter.GET("/read", func(c echo.Context) error {
		t := schema.ContentTable{}
		controller.ReadORMDataSnippet(t,mysql.OrmConn())
		return c.String(http.StatusOK,"리드")
	})

	// Update
	ormRouter.PUT("/update", func(c echo.Context) error {
		controller.UpdateORMDataSnippet(schema.ContentTable{4,"content4"},mysql.OrmConn())
		return c.String(http.StatusOK,"업데이트")
	})

	//Delete
	ormRouter.DELETE("/delete", func(c echo.Context) error {
		u := schema.ContentTable{Id:4}
		controller.DeleteORMDataSnippet(u,mysql.OrmConn())
		return c.String(http.StatusOK, "딜리트")
	})
```