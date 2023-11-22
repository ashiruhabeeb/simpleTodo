package handler

import (
	"log/slog"
	"strconv"

	"github.com/ashiruhabeeb/simpleTodoApp/entity"
	"github.com/ashiruhabeeb/simpleTodoApp/handler/request"
	"github.com/ashiruhabeeb/simpleTodoApp/logger"
	"github.com/ashiruhabeeb/simpleTodoApp/repository"
	"github.com/labstack/echo/v4"
)

type todoController struct {
	repo repository.TodoRepo
	log slog.Logger
}

func NewTodoService(repo repository.TodoRepo, log slog.Logger) todoController {
	logger := logger.NewSlogHandler()
	
	return todoController{repo: repo, log: logger}
}

func(td *todoController) Store(e echo.Context) error {
	var todorequest request.TodoRequest

	if err := e.Bind(&todorequest); err != nil {
		td.log.Error(err.Error())
		return echo.NewHTTPError(400, err.Error())
	}
	
	if errValidate := e.Validate(todorequest); errValidate != nil {
		td.log.Warn(errValidate.Error())
		return echo.NewHTTPError(400, errValidate.Error())
	}

	todo := todorequest.ToEntity()

	todoId, err := td.repo.InsertTodo(*todo)
	if err != nil {
		td.log.Error(err.Error())
		return echo.NewHTTPError(500, err.Error())
	}

	return e.JSON(201, todoId)
}

func (td *todoController) GetTodos(e echo.Context) error {
	todos, err := td.repo.GetTodos()
	if err != nil {
		td.log.Error("internal serval error: %v", err)
		return echo.NewHTTPError(500, err)
	}

	return e.JSON(200, todos)
}

func (td *todoController) GetTodo(e echo.Context) error {
	todoid := e.Param("todoid")

	todo_id, err := strconv.Atoi(todoid)
	if err != nil {
		td.log.Warn("string parse to int failure: %v", err)
		return e.JSON(400, err)
	}

	todo, err := td.repo.GetTodo(todo_id)
	if err != nil {
		td.log.Error("get todo by id failure: %v", err)
		return echo.NewHTTPError(500, err)
	}

	return e.JSON(200, todo)
}

func (td *todoController) UpdateTodo(e echo.Context) error {
	var payload struct {
		Title	string	`json:"title" validate:"required"`
		Description string `json:"description" validate:"required"`
		Start_at	string	`json:"start_at" validate:"required"`
		End_At		string `json:"end_at" validate:"required"`
	}

	todoid := e.Param("todoid")

	todo_id, err := strconv.Atoi(todoid)
	if err != nil {
		td.log.Warn("string parse to int failure: %v", err)
		return e.JSON(400, err)
	}

	if err := e.Bind(&payload); err != nil {
		td.log.Error(err.Error())
		return echo.NewHTTPError(400, err.Error())
	}

	if errValidate := e.Validate(payload); errValidate != nil {
		td.log.Warn(errValidate.Error())
		return echo.NewHTTPError(400, errValidate.Error())
	}

	if err = td.repo.UpdateTodo(todo_id, payload.Title, payload.Description, payload.Start_at, payload.End_At); err != nil {
		td.log.Error("internal server error: %v", err)
		return echo.NewHTTPError(500, err)
	}
	return e.JSON(200, "todo record updted")
}

func (td *todoController) DeleteTodo(e echo.Context) error {
	todoid := e.Param("todoid")

	todo_id, err := strconv.Atoi(todoid)
	if err != nil {
		td.log.Warn("string parse to int failure: %v", err)
		return e.JSON(400, err)
	}

	todo := entity.Todo{TodoID: todo_id}

	err = td.repo.DeleteTodo(todo.TodoID)
	if err != nil {
		td.log.Error("internal server error: %v", err)
		return echo.NewHTTPError(500, err)
	}

	return e.JSON(200, "todo deleted!")
}
