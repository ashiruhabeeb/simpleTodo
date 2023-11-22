package middleware

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ashiruhabeeb/simpleTodoApp/token"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware creates an echo middleware for authorization
func AuthMiddleware(next echo.HandlerFunc, tknMaker token.TokenMaker) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("authorization")

		if len(authHeader) == 0 {
			err := errors.New("authorization header is not provided")
			return c.JSON(401, err.Error())
		}

		fields := strings.Fields(authHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header malformed")
			return c.JSON(401, err.Error())
		}

		authHeaderType := strings.ToLower(fields[0])
		if authHeaderType != "bearer" {
			err := fmt.Errorf("unsupported authorization type %s", authHeaderType)
			return c.JSON(401, err.Error())
		}

		token := fields[1]
		payload, err := tknMaker.VerifyToken(token)
		if err != nil {
			return c.JSON(401, err.Error())
		}

		c.Set("auth_payload", payload)
		
		return next(c)
	}
}
