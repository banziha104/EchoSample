package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"EchoSample/model/schema"
	"net/http"
	"github.com/labstack/gommon/log"
)

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

func Read(db *sql.DB) schema.PaginationRespone {

	rows, err := db.Query("SELECT * FROM echoTable") // 여러 행인 경우엔 Query를 이용
	if err != nil{
		log.Fatal("database error:" ,err)
	}

	results := []schema.TestStruct{} // 결과를 받을 구조체 생성
	temp := schema.TestStruct{} // 임시 값을 담을 구조체 생성

	for rows.Next() { // 결과 값 만큼실행
		var id int
		var name string
		err = rows.Scan(&id,&name) // 데이터 리드 진행
		if err != nil {
			log.Fatal(err)
		}
		temp.Id = id
		temp.Name = name
		results =append(results,temp)
	}

	defer rows.Close() // Db를 닫음
	return schema.PaginationRespone{http.StatusOK, len(results),results}
}
