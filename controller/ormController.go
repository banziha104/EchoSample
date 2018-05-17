package controller

import (
"github.com/jinzhu/gorm"
"EchoSample/model/schema"
	"fmt"
)

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