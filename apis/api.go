package apis

import (
	"database/sql"
	_"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool `json:"completed"`
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "123456"
	DB_NAME     = "todos"
)
	var DB = ConnectDB()
	var Err error

	func ConnectDB()*sql.DB{
	 db, err := sql.Open("postgres", "user=postgres password=123456 dbname=todos sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected...")
		return db
	}
	}
func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

var todoInsert = Todo{
	Id: 12,
	Title: "Test",
	Completed: false,
}


func GetTodo(c echo.Context) error {
		sqlStatement := "SELECT * FROM todos ORDER BY id ASC "
		rows, err := DB.Query(sqlStatement)
		if err != nil {
			fmt.Println(err)
		}
		defer rows.Close()
		var result []Todo
		for rows.Next(){
			var model Todo
			if err := rows.Scan(&model.Id, &model.Title, &model.Completed);
			err !=nil{
				return err
			}
			fmt.Println(model)
			result = append(result, model )
			fmt.Println(result)
		}
		 return c.JSON(200,result)
	}

	func PostTodo(c echo.Context,data Todo) error{
		t:= Todo{
			Id: data.Id,
			Title: data.Title,
			Completed: data.Completed,
		}
		if Err:= c.Bind(t); Err!=nil{
			return Err
		}
		_, err := DB.Query(sqlStatement)
		if err != nil {
			fmt.Println(err)
		}
	}