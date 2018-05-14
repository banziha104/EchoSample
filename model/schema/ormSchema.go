package schema

type ContentTable struct{
	Id int `gorm:"primary_key"`
	Content string
}

