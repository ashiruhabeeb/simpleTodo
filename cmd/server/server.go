package server

import (
	"os"

	"github.com/ashiruhabeeb/simpleTodoApp/config"
	"github.com/ashiruhabeeb/simpleTodoApp/db"
	"github.com/ashiruhabeeb/simpleTodoApp/logger"
	"github.com/ashiruhabeeb/simpleTodoApp/router"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

type App struct {}

func(a * App) AppRun() error {
	// implement logger
	logger := logger.NewSlogHandler()

	// load env
	err := config.ENV()
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("app env variables successfully loaded")

	// start database
	db, err := db.ConnectDB(os.Getenv("DB_DSN"))
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("databse connection established")
	defer db.Close()

	// create app router
	echo := echo.New()
	echo.Use(middleware.Logger())	// middleware attached
	echo.Use(middleware.Recover())	// middleware attached

	// fetch port value from .env
	port := os.Getenv("APP_PORT")

	// setup app routes
	router.SetupRoutes(echo, port)

	return nil
}
