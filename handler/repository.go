package handler

import "github.com/labstack/echo/v4"

type TodoHandlerRepo interface {
	Store(e echo.Context) error
}