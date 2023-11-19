package repository

import (
	"database/sql"
	"fmt"

	"github.com/ashiruhabeeb/simpleTodoApp/entity"
)

// todoRepo holds db object of type database/sql package
type todoRepo struct {
	db *sql.DB
}

// NewTodoRepo constructor creates a new instance of todoRepo object
func NewTodoRepo(db *sql.DB) *todoRepo {
	return &todoRepo{db: db}
}

// InsertUser creates a new user record in the users table
func (u *todoRepo) InsertUser(t entity.Todo) (int, error) {
	err := u.db.QueryRow(insertTodo, t.Title, t.Description, t.StartAt, t.EndAt).Scan(&t.TodoID)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return t.TodoID, nil
}

