# MongoDB 연동

1. 구조체 선언 

```go
type testStruct struct {
	ID int `bson:"id"`
	Name string `bson:"name"` // 태그를 기반으로 찾음
}
```

2. 데이터베이스 컨넥션 연결

```go
session, err := mgo.Dial("127.0.0.1:27017") // 세션을 열음

	if err != nil {
		log.Fatal(err)
	}

	defer session.Close() // 끝나고 세션을 닫음

```

3. Document 조작 실행

```go
	var results []testStruct // 결과값을 만듬
	collection := session.DB("echo_test").C("echo_docs") //echo_test DB에 echo_docs Documents를 가져옮
	collection.Find(nil).All(&results) // 모든 데이터를 가져옮
	for _,result := range results{ // 결과값을 가져옮
		fmt.Println(result)
	}
```