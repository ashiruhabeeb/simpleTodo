package handler

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/ashiruhabeeb/simpleTodoApp/entity"
	"github.com/ashiruhabeeb/simpleTodoApp/logger"
	"github.com/ashiruhabeeb/simpleTodoApp/repository"
	"github.com/ashiruhabeeb/simpleTodoApp/token"
	"github.com/ashiruhabeeb/simpleTodoApp/utils"
	"github.com/labstack/echo/v4"
)

// Holds userController objects
type userController struct {
	repo  repository.UserRepo
	log	  slog.Logger
	token token.TokenMaker
}

// Initialize an instance of userController
func NewUserController(repo repository.UserRepo, log slog.Logger, token token.TokenMaker) userController {
	logger :=logger.NewSlogHandler()

	return userController{repo: repo, log: logger, token: token}
}

// SignUp creates a new user record in the users table
func (uc *userController) SignUp(e echo.Context) error {
	var payload struct {
		Username		string	`json:"username" validate:"required;min=2,max=225"`
		FullName		string	`json:"fullname" validate:"required;min=2,max=225"`
		Email			string	`json:"e-mail" validate:"required,email"`
		Password		string	`json:"password" validate:"required,min=8"`
		ConfirmPasword	string	`json:"confirm_password" validate:"required"`
		Phone			string	`json:"phone" validate:"required,e164"`
	}

	if err := e.Bind(&payload); err != nil {
		uc.log.Error(err.Error())
		return HandlerError(e, 400, err)
	}

	if payload.ConfirmPasword != payload.Password {
		uc.log.Warn("password mismatch! Ensure ConfirmPazzword matches Pasword")
		return HandlerError(e, 400, nil)
	}

	err := e.Validate(payload)
	if err != nil {
		uc.log.Warn(err.Error())
		return HandlerError(e, 400, err)
	}

	payload.Password, err = utils.HashPWD(payload.Password)
	if err != nil {
		uc.log.Warn("password hash generation failure:", err)
		return HandlerError(e, 500, err)
	}

	entity := &entity.User{
		Username:  payload.Username,
		FullName:  payload.FullName,
		Email:     payload.Email,
		Password:  payload.Password,
		Phone:     payload.Phone,
		CreatedAt: time.Now().Local().String(),
	}

	id, err := uc.repo.InsertUser(*entity)
	if err != nil {
		uc.log.Error("user create failure: %v", err)
		return HandlerError(e, 500, err)
	}

	accesToken, _, err := uc.token.GenerateAccessToken(entity.Username, time.Minute * 20)
	if err != nil {
		uc.log.Error("access token generation failure: %v", err)
		return HandlerError(e, 500, err)
	}

	refreshToken, _, err := uc.token.GenerateRefreshToken(entity.Username, time.Hour * 24 * 7 * 51)
	if err != nil {
		uc.log.Error("refresh token generation failure: %v", err)
		return HandlerError(e, 500, err)
	}

	e.SetCookie(&http.Cookie{
		Name:       "Authorization",
		Value:      accesToken,
		Path:       "/",
		MaxAge:     int(time.Now().Add(time.Minute * 20).Unix()),
		Secure:     true,
		HttpOnly:   true,
		SameSite:   http.SameSiteDefaultMode,
	})

	e.SetCookie(&http.Cookie{
		Name:       "Authorization",
		Value:      refreshToken,
		Path:       "/",
		MaxAge:     int(time.Now().Add(time.Hour * 24 * 7 * 51).Unix()),
		Secure:     true,
		HttpOnly:   true,
		SameSite:   http.SameSiteDefaultMode,
	})

	return e.JSON(201, id)
}
