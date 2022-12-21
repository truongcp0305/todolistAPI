package main

import (
	"Go/apis"
	"fmt"
	"net/http"
	_ "net/http"

	"github.com/labstack/echo"
)
type Todo struct {
	Id string `json:"id"`
}

func main() {
	var DB = apis.ConnectDB()
	e := echo.New()
	e.GET("/todos", apis.GetTodo)
	e.GET("/", apis.Hello)
	e.POST("todos/add", func(c echo.Context) error{
		 req := new(Todo)
		if model := c.Bind(&req); model != nil{
			//return c.String(http.StatusBadRequest, "bad request")
		}
		fmt.Println(&req)
		fmt.Println(req)
		statement := "INSERT INTO test (Id) VALUES ($1)"
		res, err := DB.Query(statement, &req.Id)
		if err!=nil{
			fmt.Println(err)
		}else{
			fmt.Println(res)
			return c.JSON(http.StatusCreated, req)
		}
		return c.String(200,"ok")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
