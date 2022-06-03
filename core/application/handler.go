package app_todo

// all handlers are called in /server/routes.go

import (
	"context"
	"database/sql"
	"hex-arch-go/core/entities"
	todo "hex-arch-go/core/infrastructure"
	"hex-arch-go/internal/helpers"

	"github.com/labstack/echo/v4"
)

// The HTTP Handler for TODO
type ToDoHTTPService struct {
	gtw todo.ToDoGateway
}

func (t *ToDoHTTPService) ListHandler(c echo.Context) error {
	status, res := t.gtw.ListToDo()
	return c.JSON(status, res)
}

func (t *ToDoHTTPService) CreateHandler(c echo.Context) error {
	// New entitie of ToDo
	td := new(entities.ToDo)
	// Bind the data into struct
	c.Bind(td)
	// if we have a validation of data we need do it here,
	// and for example, return an http.StatusBadRequest error when have a empty field

	// in this case, now i create the possible new ID of user for not waiting of infrastructure response
	td.ID = helpers.UUID()
	// or i can get it from the frontend / another micro service

	// we need the complete critical data into the entity before call the gateway
	status, res := t.gtw.CreateToDo(td)
	return c.JSON(status, res)
}

// Constructor
func NewToDoHTTPService(ctx context.Context, db *sql.DB) *ToDoHTTPService {
	return &ToDoHTTPService{todo.NewToDoGateway(ctx, db)}
}
