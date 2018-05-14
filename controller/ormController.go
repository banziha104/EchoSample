package controller

import (
"github.com/jinzhu/gorm"
"EchoSample/model/schema"
	"fmt"
)

func CreateORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	err := db.Create(c).Error // insert into content_tables
	if err != nil {
		panic(err.Error())
	}
}


func ReadORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	result := db.Find(&c) // Select * from content_tables
	fmt.Println(result.Value)
}


func UpdateORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	//result := db.Model(&c).Update(4,"ddd")
	result := db.Model(&c).Updates(schema.ContentTable{Id:4,Content:"ddd"})
	fmt.Println(result.Value)
}


func DeleteORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	result := db.Delete(&c)
	fmt.Println(result)
}