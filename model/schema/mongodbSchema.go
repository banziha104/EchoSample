package schema

type NosqlStruct struct {
	ID int `bson:"id"`
	Name string `bson:"name"` // 태그를 기반으로 찾음
}
