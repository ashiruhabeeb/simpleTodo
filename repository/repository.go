package repository

import "github.com/ashiruhabeeb/simpleTodoApp/entity"

// Implements todoRepo methods
type TodoRepo interface {
	InsertUser(t entity.Todo) (int, error)
}
