package controller

import (
"github.com/jinzhu/gorm"
"EchoSample/model/schema"
)

func CreateORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	err := db.Create(c).Error
	if err != nil {
		panic(err.Error())
	}
}


func ReadORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	err := db.Where("").Error
	if err != nil {
		panic(err.Error())
	}
}


func UpdateORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	err := db.Create(c).Error
	if err != nil {
		panic(err.Error())
	}
}


func DeleteORMDataSnippet(c schema.ContentTable,db *gorm.DB){
	err := db.Create(c).Error
	if err != nil {
		panic(err.Error())
	}
}