package todo

import (
	"hex-arch-go/core/entities"
)

// Mocked repository
type ToDoMock struct{}

func (t *ToDoMock) insertToDoInDb(td *entities.ToDo) {}
func (t *ToDoMock) listToDoInDb() []entities.ToDo {
	return []entities.ToDo{
		{
			ID: "00a95ff5-496c-4c72-87b8-fbc787a53aae",
			To: "To",
			Do: "Do!",
		},
		{
			ID: "e2663c89-e637-445d-aff2-70836d26aee5",
			To: "Too",
			Do: "Doo!",
		},
		{
			ID: "d1826c77-ef4d-47fb-b1a9-24fb0df8551e",
			To: "Tooo",
			Do: "Dooo!",
		},
	}
}

// Constructor
func NewToDoMockStorage() ToDoStorage {
	return &ToDoMock{}
}
