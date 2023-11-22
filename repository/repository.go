package repository

import "github.com/ashiruhabeeb/simpleTodoApp/entity"

// Implements todoRepo methods
type TodoRepo interface {
	InsertTodo(t entity.Todo) (int, error)
	GetTodos()([]entity.Todo, error)
	GetTodo(todo_id int)(*entity.Todo, error)
	UpdateTodo(todo_id int, title, description, start_at, end_at string) error
	DeleteTodo(todo_id int) error
}

// Implements userRepo methods
type UserRepo interface {
	InsertUser(u entity.User)(int, error)
	GetUser(user_id int)(*entity.User, error)
	GetUsers()([]entity.User, error)
	UpdateUser(user_id int, username, fullname, address, dob string) error
	DeleteUser(user_id int) error
}
