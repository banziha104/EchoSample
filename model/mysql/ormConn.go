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