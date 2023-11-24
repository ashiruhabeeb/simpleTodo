package handler

import "github.com/labstack/echo/v4"

func HandlerError(c echo.Context, code int, err error) error {
	return c.JSON(
		code,
		map[string]any{"err": err},
	)
}
