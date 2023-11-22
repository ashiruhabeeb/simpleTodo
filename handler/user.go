package handler

import (
	"log/slog"

	"github.com/ashiruhabeeb/simpleTodoApp/logger"
	"github.com/ashiruhabeeb/simpleTodoApp/repository"
	"github.com/labstack/echo/v4"
)

type userController struct {
	repo repository.UserRepo
	log	 slog.Logger
}

func NewUserController(repo repository.UserRepo, log slog.Logger) userController {
	logger :=logger.NewSlogHandler()

	return userController{repo: repo, log: logger}
}

func (uc *userController) SignUp(e echo.Context) error {


	return nil
}