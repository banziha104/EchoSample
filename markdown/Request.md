# 요청처리

- QueryParams() : get param 받기

```go
e.GET("/", func(c echo.Context) error {
		results := mysql.Read(mysql.DBConn())
		fmt.Println(c.QueryParams()) // 파라미터를 맵으로 받아옮
		fmt.Println(c.QueryParams("name")) // 파라미터를 찾아 리턴
		return c.JSONPretty(http.StatusOK,results,"    ")
	})
```

- FormParams(), FormValue(target) : post body 파싱

```go
	e.POST("/", func(c echo.Context) error {
		fmt.Println(c.FormParams()) // 파라미터 전체를 받음
		fmt.Println(c.FormValue("name")) // 파라미터를 찾아 리턴
		return c.String(http.StatusOK, "postResponse")
	})
```

- Param() : path 파싱 

```go

	e.POST("/:id", func(c echo.Context) error {
		v := c.Param("id")
		fmt.Println(v)
		return c.String(http.StatusOK,v)
	})
```