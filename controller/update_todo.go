package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

type UpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewUpdateTodoController(e *echo.Echo, db *sql.DB) {
	e.PATCH("/todos/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")

		var request UpdateRequest

		err := json.NewDecoder(ctx.Request().Body).Decode(&request)

		if err != nil {
			return ctx.String(http.StatusBadRequest, "Invalid request body")
		}

		if request.Title == "" && request.Description == "" {
			return ctx.String(http.StatusBadRequest, "No fields to update")
		}

		query := "UPDATE todos SET"
		var params []interface{}

		if request.Title != "" {
			query += " title = ?,"
			params = append(params, request.Title)
		}

		if request.Description != "" {
			query += " description = ?,"
			params = append(params, request.Description)
		}

		query = strings.TrimSuffix(query, ",")

		query += " WHERE id = ?"
		params = append(params, id)

		_, err = db.Exec(query, params...)

		if err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusOK, "OK")
	})

}
