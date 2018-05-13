# 라우팅

- /users/:id : 파라미터에서 파싱
- /users/new : 정적 매칭
- /users/* : users 다음에 무엇이와도 가능

### Group

```go

g := e.Group("/api") // /api/ 에 옮
g.GET("/", func(c echo.Context) error {
    return c.String(http.StatusOK,"/api/getResponse")
})

```