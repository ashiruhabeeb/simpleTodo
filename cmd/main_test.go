package main

import (
	"os"
	"testing"

	"github.com/ashiruhabeeb/simpleTodoApp/db"
	"github.com/ashiruhabeeb/simpleTodoApp/logger"
	"github.com/ashiruhabeeb/simpleTodoApp/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func TestAppRun(t *testing.T) {
	// implement logger
	logger := logger.NewSlogHandler()
	
	// load env
	godotenv.Load()

	// start database
	db, err := db.ConnectDB(os.Getenv("DB_TEST"))
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("✅ databse connection established")

	// create app router
	echo := echo.New()
	echo.Use(middleware.Logger())	// middleware attached
	echo.Use(middleware.Recover())	// middleware attached

	// fetch port value from .env
	port := os.Getenv("APP_PORT")

	// setup app routes
	router.SetupRoutes(echo, port, db)
}

func TestMain(m *testing.M){

	os.Exit(m.Run())
}
