package todo

// This is a integration test

import (
	"context"
	"database/sql"
	"fmt"
	"hex-arch-go/core/entities"
	"hex-arch-go/env"
	"hex-arch-go/internal/database"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	ctx     = context.Background()
	db      = initTestDB()
	storage = NewToDoStorage(ctx, db)
)

const (
	MAX_NUMBER_OF_TODO_INSERTED = 10
)

func TestInsertToDoInDb(t *testing.T) {
	testCases := []struct {
		Name     string
		ToDoData entities.ToDo
	}{
		{
			Name: "It should create a new ToDo",
			ToDoData: entities.ToDo{
				To: "To!",
				Do: "Do!",
			},
		},
		{
			Name: "It should create a new ToDo",
			ToDoData: entities.ToDo{
				To: "ToO!",
				Do: "DoO!",
			},
		},
		{
			Name: "It should create a new ToDo",
			ToDoData: entities.ToDo{
				To: "ToOo!",
				Do: "DoOo!",
			},
		},
		{
			Name: "It should create a new ToDo",
			ToDoData: entities.ToDo{
				To: "ToX",
				Do: "DoX",
			},
		},
		// More cases, it depends of your logic implemented on Storage
	}

	for i := range testCases {
		tc := testCases[i]

		// i recommend do not use t.Run() because the context in test is another,
		// and is possible what we want debug and check errors in the process

		t.Log(tc.Name)

		// Create the TODO
		tc.ToDoData.ID = uuid.NewString()
		storage.insertToDoInDb(&tc.ToDoData)

		// Check the TODO created
		var td entities.ToDo
		querystring := `SELECT id, _to, _do FROM todo WHERE id = $1 LIMIT 1`
		row := db.QueryRowContext(ctx, querystring, tc.ToDoData.ID)
		err := row.Scan(&td.ID, &td.To, &td.Do)

		// Asserts
		assert.NoError(t, err)
		assert.Equal(t, tc.ToDoData, td)
	}

	// Clean the table after test
	db.ExecContext(ctx, "TRUNCATE todo")
}

func TestListToDoInDb(t *testing.T) {
	// Create a generic list of Todos
	var todos []entities.ToDo
	for i := 0; i < MAX_NUMBER_OF_TODO_INSERTED; i++ {
		var td entities.ToDo
		td.ID = uuid.NewString()
		td.To = fmt.Sprintf("To %d!", i)
		td.Do = fmt.Sprintf("Do %d!", i)
		todos = append(todos, td)
	}

	testCases := []struct {
		Name     string
		Expected []entities.ToDo
	}{
		{
			Name:     "It should be empty list of TODO",
			Expected: nil,
		},
		{
			Name:     fmt.Sprintf("It should be a list of %d TODO", len(todos)),
			Expected: todos,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		// i recommend do not use t.Run() because the context in test is another,
		// and is possible what we want debug and check errors in the process

		t.Log(tc.Name)

		// Insert if i expect results
		if tc.Expected != nil {
			for x := range tc.Expected {
				querystring := `INSERT INTO todo (id, _to, _do) VALUES ($1, $2, $3)`
				_, _ = db.ExecContext(ctx, querystring, tc.Expected[x].ID, tc.Expected[x].To, tc.Expected[x].Do)
			}
		}

		// Get list
		td := storage.listToDoInDb()

		// Asserts
		assert.Equal(t, tc.Expected, td)
	}

	// Clean the table after test
	db.ExecContext(ctx, "TRUNCATE todo")
}

func TestNewToDoStorage(t *testing.T) {
	as := NewToDoStorage(ctx, nil)
	var expect ToDoStorage = &ToDoService{db: nil, ctx: ctx}
	assert.Equal(t, as, expect)
}

// Just for close the DB connection open in the test
func TestClean(t *testing.T) {
	db.Close()
}

// We need start the database, but we can do in another database with the same schema
func initTestDB() *sql.DB {
	// The .env config for testing
	_env := env.GetEnv("../../.env.test")
	return database.NewCockRoachDatabase(ctx, _env).ConnectDB()
}
