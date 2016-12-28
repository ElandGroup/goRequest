package main

import (
	"goRequest/logHelper"
	"net/http"

	"github.com/labstack/echo"
)

var (
	e = echo.New()
)

func main() {

	e.GET("/", func(c echo.Context) error {
		logHelper.Warning.Println("hello good")
		return c.String(http.StatusOK, "Hello echo api")
	})
	RegisterUserApi()
	logHelper.Warning.Println("RegisterUserApi end")

	e.Start(":1111")
}
