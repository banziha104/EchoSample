package main

import (
	"EchoSample/router"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	router.APIRouter(e)
	e.Logger.Fatal(e.Start(":8081"))
}