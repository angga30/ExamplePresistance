package main

import (
	"net/http"
	a "prod/lovalf"
	"github.com/labstack/echo"
)

func main() {
	var Con a.Connection
	e := echo.New()
	Con = a.SetDb()
	az := Con.GetDb().Db
	
	er := az.Ping()
	if er != nil {
		print(er.Error())
	}
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	e.Logger.Fatal(e.Start(":1323"))
}

