# MySql 연동

- 설치 : go get -u github.com/go-sql-driver/mysql

1. 받을 데이터의 구조체를 생성

```go
type testStruct struct { // DB 아키텍쳐 구조체
	id int
	name string
}
```

2. 데이터베이스와 연결

```go
func DBConn() (db *sql.DB){ // DB 연결부
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "echoTest"
	db, err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		panic(err.Error())
	}
	return db
}
```

3. sql문 작성

```go
func Read(db *sql.DB){

	rows, err := db.Query("SELECT * FROM echoTable") // 여러 행인 경우엔 Query를 이용
	if err != nil{
		log.Fatal("database error:" ,err)
	}

	results := []testStruct{} // 결과를 받을 구조체 생성
	temp := testStruct{} // 임시 값을 담을 구조체 생성

	for rows.Next() { // 결과 값 만큼실행
		var id int
		var name string
		err = rows.Scan(&id,&name) // 데이터 리드 진행
		if err != nil {
			log.Fatal(err)
		}
		temp.id = id
		temp.name = name
		results =append(results,temp)
	}
	fmt.Println(results)
	defer rows.Close() // Db를 닫음
}

```

### 전체소스

```go
package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/labstack/gommon/log"
	"fmt"
)
type testStruct struct { // DB 아키텍쳐 구조체
	Id int          `json:"id"`
	Name string     `json:"name"`
}
func DBConn() (db *sql.DB){ // DB 연결부
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "1234"
	dbName := "echoTest"
	db, err := sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	if err != nil{
		panic(err.Error())
	}
	return db
}

func Read(db *sql.DB){

	rows, err := db.Query("SELECT * FROM echoTable") // 여러 행인 경우엔 Query를 이용
	if err != nil{
		log.Fatal("database error:" ,err)
	}

	results := []testStruct{} // 결과를 받을 구조체 생성
	temp := testStruct{} // 임시 값을 담을 구조체 생성

	for rows.Next() { // 결과 값 만큼실행
		var id int
		var name string
		err = rows.Scan(&id,&name) // 데이터 리드 진행
		if err != nil {
			log.Fatal(err)
		}
		temp.id = id
		temp.name = name
		results =append(results,temp)
	}
	fmt.Println(results)
	defer rows.Close() // Db를 닫음
}

```