# 라우팅

- /users/:id : 파라미터에서 파싱
- /users/new : 정적 매칭
- /users/* : users 다음에 무엇이와도 가능

### Group

> 동일한 하이어라키를 그룹핑

```go

g := e.Group("/api") // /api
g.GET("/", func(c echo.Context) error { // /api/
    return c.String(http.StatusOK,"/api/getResponse")
})

```

### URI 빌딩

> 핸들러와 라우터를 이용해 구현

```go
h := func(c echo.Context) error {
		return c.String(http.StatusOK,"OK")
	}

	e.GET("/handler",h)
```