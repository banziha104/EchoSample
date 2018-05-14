package router

import (
	"github.com/labstack/echo"
	"net/http"
	"EchoSample/model/mysql"
	"fmt"
	"EchoSample/controller"
	"EchoSample/model/schema"
)


func APIRouter(e *echo.Echo){
	g := e.Group("/api")
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"/api/ : group Response")
	})


	e.GET("/", func(c echo.Context) error {
		results := mysql.Read(mysql.DBConn())
		fmt.Println(c.QueryParams())
		return c.JSONPretty(http.StatusOK,results,"    ")
	})

	e.POST("/", func(c echo.Context) error {
		fmt.Println(c.FormParams())
		fmt.Println(c.FormValue("name")) // 파라미터를 찾아 리턴
		return c.String(http.StatusOK, "postResponse")
	})

	e.POST("/:id", func(c echo.Context) error {
		v := c.Param("id")
		fmt.Println(v)
		return c.String(http.StatusOK,v)
	})


	ormRouter := e.Group("/orm")

	ormRouter.POST("/insert", func(c echo.Context) error {
		controller.CreateORMDataSnippet(schema.ContentTable{4,"content4"},mysql.OrmConn())
		return c.String(http.StatusOK,"인설또")
	})

	//ormRouter.GET("/read", func(c echo.Context) error {
	//
	//})
	//
	//ormRouter.PUT("/read/:id", func(c echo.Context) error {
	//
	//})
	//
	//ormRouter.DELETE("/read/:id", func(c echo.Context) error {
	//
	//})
}