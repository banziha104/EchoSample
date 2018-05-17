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

	//localhost:8081/ get요청
	e.GET("/", func(c echo.Context) error {
		results := mysql.Read(mysql.DBConn())
		fmt.Println(c.QueryParams())
		return c.JSONPretty(http.StatusOK,results,"    ")
	})
	//localhost:8081/ post요청
	e.POST("/", func(c echo.Context) error {
		fmt.Println(c.FormParams())
		fmt.Println(c.FormValue("name")) // 파라미터를 찾아 리턴
		return c.String(http.StatusOK, "postResponse")
	})

	e.PUT("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"방금 만든 풋")
	})

	e.POST("/:id", func(c echo.Context) error {
		v := c.Param("id")
		fmt.Println(v)
		return c.String(http.StatusOK,v)
	})


	//localhost8081:/api/
	g := e.Group("/api")
	g.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK,"/api/ : group Response")
	})

	//www.naver.com/orm/
	ormRouter := e.Group("/orm")

	// Read
	ormRouter.GET("/", func(c echo.Context) error {
		t := schema.ContentTable{}
		controller.ReadORMDataSnippet(t,mysql.OrmConn())
		return c.String(http.StatusOK,"리드")
	})

	// Create
	ormRouter.POST("/", func(c echo.Context) error {
		controller.CreateORMDataSnippet(schema.ContentTable{4,"content4"},mysql.OrmConn())
		return c.String(http.StatusOK,"인설또")
	})


	// Update
	ormRouter.PUT("/", func(c echo.Context) error {
		controller.UpdateORMDataSnippet(schema.ContentTable{4,"content4"},mysql.OrmConn())
		return c.String(http.StatusOK,"업데이트")
	})

	//Delete
	ormRouter.DELETE("/", func(c echo.Context) error {
		u := schema.ContentTable{Id:4}
		controller.DeleteORMDataSnippet(u,mysql.OrmConn())
		return c.String(http.StatusOK, "딜리트")
	})
}

/***
http://apis.data.go.kr/1192000/openapi/service/ManageAcst0400Service/getAcst0400List?
serviceKey=i9opnT0CNWj0dfjeUmoProOy3c%2BqZNdfztvalVl624EISpMpkXLDvVzwuuA8n8BnYnMqOjKlZIoBQLm%2FpX%2Fyqg%3D%3D&
numOfRows=10&
pageSize=10&
pageNo=1&
startPage=1&
fromDt=20151001&toDt=20151003
 */