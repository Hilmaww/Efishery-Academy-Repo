package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type M map[string]interface{}

func main() {
	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Hello from /index"
		return ctx.String(http.StatusOK, data)
	})
	r.GET("/users/:nama", func(ctx echo.Context) error {
		nama := ctx.Param("nama")
		data := M{"Success": true, "data": M{"name": nama}}
		return ctx.JSON(http.StatusOK, data)
	})
	r.Start(":9000")
}
