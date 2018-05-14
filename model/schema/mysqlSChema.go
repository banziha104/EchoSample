package schema

import "github.com/jinzhu/gorm"

type User struct {
	Name string
	Email string
}

type PaginationRespone struct {
	StatusCode int `json:"statusCode"`
	Count int `json:"count"`
	Results []TestStruct `json :"results"`
}

type TestStruct struct { // DB 아키텍쳐 구조체
	Id int          `json:"id"`
	Name string     `json:"name"`
}

type GormTestStruct struct {
	gorm.Model
	TestStruct []TestStruct
}