package request

import "github.com/ashiruhabeeb/simpleTodoApp/entity"

type TodoRequest struct {
	Title		string	`json:"title" validate:"required"`
	Description string	`json:"description" validate:"required"`
	StartAt		string	`json:"start_at" validate:"required"`
	EndAt		string	`json:"end_at" validate:"required"`
}

func(todoRequest *TodoRequest) ToEntity() *entity.Todo {
	return &entity.Todo{
		Title:       todoRequest.Title,
		Description: todoRequest.Description,
		StartAt:     todoRequest.StartAt,
		EndAt:       todoRequest.EndAt,
	}
}
