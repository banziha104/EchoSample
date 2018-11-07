# 에러 핸들링

- 기본 에러 설정

```go
e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
  return func(c echo.Context) error {
    // Extract the credentials from HTTP request header and perform a security
    // check

    // For invalid credentials
    return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

    // For valid credentials call next
    // return next(c)
  }
})
```

- 커스텀 에러

```go
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	c.Logger().Error(err)
}

e.HTTPErrorHandler = customHTTPErrorHandler
```