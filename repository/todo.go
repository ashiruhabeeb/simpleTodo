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
func (tr *todoRepo) InsertUser(t entity.Todo) (int, error) {
	err := tr.db.QueryRow(insertTodo, t.Title, t.Description, t.StartAt, t.EndAt).Scan(&t.TodoID)
	if err != nil {
		return 0, fmt.Errorf(err.Error())
	}
	return t.TodoID, nil
}

// GetTodos returns all todo records in the todo table
func (tr *todoRepo) GetTodos()([]entity.Todo, error){
	rows, err := tr.db.Query(getTodos)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := []entity.Todo{}
	for rows.Next() {
		var todo entity.Todo
		err := rows.Scan(&todo.TodoID, &todo.Title, &todo.Description, &todo.Completed, &todo.StartAt, &todo.EndAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}

// GetTodo returns a single todo record from the todso table
func (tr *todoRepo) GetTodo(todo_id int)(*entity.Todo, error){
	todo := entity.Todo{}

	row := tr.db.QueryRow(getTodo, todo_id)

	err := row.Scan(&todo.TodoID, &todo.Title, &todo.Description, &todo.Completed, &todo.StartAt, &todo.EndAt)
	if err != nil {
		if err == row.Err() {
			return nil, err
		}
		return nil, err
	}
	return &todo, nil
}

// UpdateTodo updates a single todo record in the todo table
func (tr *todoRepo) UpdateTodo(todo_id int, title, description, start_at, end_at string) error {
	_, err := tr.db.Exec(updateTodo, todo_id, title, description, start_at, end_at)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo deletes a single todo record from the todo table
func (tr *todoRepo) DeleteTodo(todo_id int) error {
	_, err := tr.db.Exec(deleteTodo, todo_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("no record associated with todoID parameter: %v", err)
		}
		return err
	}
	return nil
}
