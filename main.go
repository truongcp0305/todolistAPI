package main

import (
	_"net/http"
	"github.com/labstack/echo/v4"
	"Go/apis"
)


func main() {
	e := echo.New()
	e.GET("/todos", apis.GetTodo)
	e.GET("/", apis.Hello)
	e.Logger.Fatal(e.Start(":1323"))
}
