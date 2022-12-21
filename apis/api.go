package apis

import (
	"database/sql"
	_"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
)

type Todo struct {
	Id string `json:"id"`
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
			if err := rows.Scan(&model.Id);
			err !=nil{
				return err
			}
			result = append(result, model )
		}
		 return c.JSON(200,result)
	}

	func PostTodo(c echo.Context) error{
		 req := new(Todo)
		
		 err:= c.Bind(req)
		// if err:= c.Bind(req); err!=nil{
		// 	fmt.Println(err)
		// 	fmt.Println(req)
		// 	return err
		// }	
		fmt.Println(err)
		fmt.Println(c)
		fmt.Println(req)
		// sqlStatement := "INSERT INTO todos (id, title, completed) VALUES ($1, $2, $3)"
		// _, err := DB.Query(sqlStatement, req.Id, req.Title, req.Complete)
		// if err != nil {
		// 	fmt.Println(err)
		// }else{
		// 	return c.JSON(http.StatusCreated, req)
		// }
		return c.String(http.StatusOK, "ok")
	}
