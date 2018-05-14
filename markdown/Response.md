# JSON 리스폰

```go
func APIRouter(e *echo.Echo){
	e.GET("/", func(c echo.Context) error {
        		results := mysql.Read(mysql.DBConn())
        		fmt.Println(results)
        		return c.JSON(http.StatusOK,results) // 경량화 된 전송
        	})
	
	e.GET("/", func(c echo.Context) error {
    		results := mysql.Read(mysql.DBConn())
    		fmt.Println(results)
    		return c.JSONPretty(http.StatusOK,results,"    ") // 마지막 변수는 얼마나 여백을 줄지 indent
    	})
}
```