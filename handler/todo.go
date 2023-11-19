package handler

import (
	"log/slog"

	"github.com/ashiruhabeeb/simpleTodoApp/handler/request"
	"github.com/ashiruhabeeb/simpleTodoApp/logger"
	"github.com/ashiruhabeeb/simpleTodoApp/repository"
	"github.com/ashiruhabeeb/simpleTodoApp/validator"
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
	
	if err := validator.Validate(todorequest); err != nil {
		td.log.Warn(err.Error())
		return echo.NewHTTPError(400, err.Error())
	}

	todo := todorequest.ToEntity()

	todoId, err := td.repo.InsertUser(*todo)
	if err != nil {
		td.log.Error(err.Error())
		return echo.NewHTTPError(500, err.Error())
	}

	return e.JSON(201, todoId)
}
