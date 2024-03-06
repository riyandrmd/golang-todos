package main

import (
	"golang-todos/controller"
	"golang-todos/database"

	"github.com/labstack/echo"
)

func main() {
	db := database.InitDb()
	defer db.Close()

	err := db.Ping()
	if err != nil {
		panic(err)
	}

	e := echo.New()

	controller.NewGetALLTodosController(e, db)
	controller.NewCreateTodoContoller(e, db)
	controller.NewUpdateTodoController(e, db)
	controller.NewDeleteTodoController(e, db)
	controller.NewCheckTodoController(e, db)

	e.Start(":8080")
}
