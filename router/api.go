package router

import (
	"github.com/labstack/echo"
	"net/http"
)

func APIRouter(e *echo.Echo){
	g := e.Group("/api")
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"/api/ : group Response")
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"respone")
	})
	e.POST("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "postResponse")
	})
}