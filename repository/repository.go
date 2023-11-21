package repository

import "github.com/ashiruhabeeb/simpleTodoApp/entity"

// Implements todoRepo methods
type TodoRepo interface {
	InsertUser(t entity.Todo) (int, error)
	GetTodos()([]entity.Todo, error)
	GetTodo(todo_id int)(*entity.Todo, error)
	UpdateTodo(todo_id int, title, description, start_at, end_at string) error
	DeleteTodo(todo_id int) error

}
