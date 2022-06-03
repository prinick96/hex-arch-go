package todo

import (
	"context"
	"database/sql"
	"hex-arch-go/core/entities"

	"github.com/prinick96/elog"
)

// The TODO Repository
type ToDoStorage interface {
	insertToDoInDb(td *entities.ToDo)
	listToDoInDb() []entities.ToDo
}

// The TODO Service
type ToDoService struct {
	db  *sql.DB
	ctx context.Context
}

// Get todo lists
func (t *ToDoService) listToDoInDb() []entities.ToDo {
	querystring := `SELECT id, _to, _do FROM todo ORDER BY _to ASC`
	rows, err := t.db.QueryContext(t.ctx, querystring)

	// If exist an error, return nil list
	if err != nil {
		return nil
	}

	var todos []entities.ToDo
	for rows.Next() {
		var td entities.ToDo
		err = rows.Scan(&td.ID, &td.To, &td.Do)
		go elog.New(elog.ERROR, "Error getting the list of TODO", err)

		todos = append(todos, td)
	}
	defer rows.Close()

	return todos
}

// Insert a new todo
func (t *ToDoService) insertToDoInDb(td *entities.ToDo) {
	querystring := `INSERT INTO todo (id, _to, _do) VALUES ($1, $2, $3)`
	_, err := t.db.ExecContext(t.ctx, querystring, td.ID, td.To, td.Do)
	go elog.New(elog.ERROR, "Error inserting a TODO element in Database", err)
}

// Constructor
func NewToDoStorage(ctx context.Context, db *sql.DB) ToDoStorage {
	return &ToDoService{db: db, ctx: ctx}
}
