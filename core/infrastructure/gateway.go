package todo

import (
	"context"
	"database/sql"
	"hex-arch-go/core/entities"
	"net/http"
)

// the Gateway for access to Storage
type ToDoGateway interface {
	CreateToDo(td *entities.ToDo) (int, entities.Response)
	ListToDo() (int, []entities.ToDo)
}

// The Domain Logic
type ToDoLogic struct {
	St ToDoStorage
}

// List ToDo
func (t *ToDoLogic) ListToDo() (int, []entities.ToDo) {
	// Domain logic
	return http.StatusOK, t.St.listToDoInDb()
}

// Create ToDo
func (t *ToDoLogic) CreateToDo(td *entities.ToDo) (int, entities.Response) {
	// Example domain logic
	if td.Do != "Do" {
		// Just Ok, just error message for front-end / micro-service
		return http.StatusOK, entities.Response{
			Message: "The do is invalid.",
			Success: false,
		}
	}

	// If all is ok, we can create the ToDo
	// I can use a goroutine if the response do not need anything from the infrastructure
	go t.St.insertToDoInDb(td)

	// just make a accepted response
	return http.StatusAccepted, entities.Response{
		Message: "TODO successfully added.",
		Success: true,
	}
}

// Constructor
func NewToDoGateway(ctx context.Context, db *sql.DB) ToDoGateway {
	return &ToDoLogic{NewToDoStorage(ctx, db)}
}
