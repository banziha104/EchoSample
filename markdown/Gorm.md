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
func CreateORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	// insert into content_tables
	result := db.Create(c)
	fmt.Println(result)
}


func ReadORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	// Select * from content_tables
	result := db.Find(&c)
	fmt.Println(result.Value)
}


func UpdateORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	result := db.Model(&c).Updates(schema.ContentTable{Id:4,Content:"13123"})
	fmt.Println(result.Value)
}


func DeleteORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	result := db.Delete(&c)
	fmt.Println(result)
}
```

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

- Model 태그 사용법 :

```go
type User struct {
    gorm.Model
    Birthday     time.Time
    Age          int
    Name         string  `gorm:"size:255"` // 사이즈를 255로 제한
    Num          int     `gorm:"AUTO_INCREMENT"` //자동증가

    CreditCard        CreditCard      // 1:1 관계
    Emails            []Email         // 1:N관계
    IgnoreMe          int `gorm:"-"`   // 이부분은무시
    Languages         []Language `gorm:"many2many:user_languages;"` // N:M 관계 정의
}

type Email struct {
    ID      int
    UserID  int     `gorm:"index"` // 
    Email   string  `gorm:"type:varchar(100);unique_index"` // unique 제한
    Subscribed bool
}

type Address struct {
    ID       int
    Address1 string         `gorm:"not null;unique"` // not null과 Unique 설정
    Address2 string         `gorm:"type:varchar(100);unique"`
    Post     sql.NullString `gorm:"not null"`
}

type Language struct {
    ID   int
    Name string `gorm:"index:idx_name_code"` // Create index with name, and will create combined index if find other fields defined same name
    Code string `gorm:"index:idx_name_code"` // `unique_index` also works
}

type CreditCard struct {
    gorm.Model
    UserID  uint
    Number  string
}
```