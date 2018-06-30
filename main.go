package main

import (
	"github.com/labstack/echo"
	"EchoSample/router"
)

func main() {
	e := echo.New()
	router.APIRouter(e)
	e.Logger.Fatal(e.Start(":8081"))
}