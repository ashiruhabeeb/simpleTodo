package handler

import "github.com/labstack/echo/v4"


func SomeErrorHandler(c echo.Context, code int, err error) error {
	return c.JSON(
		code,
		map[string]any{"err": err.Error()},
	)
}
