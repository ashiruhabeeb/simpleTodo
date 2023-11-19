package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	APP_ENV		string
	APP_PORT 	string

	Psql_DSN 	string
}

// ENV() loads environment variables from the .env file
func ENV() error {
	goENV := os.Getenv("APP_ENVIRONMENT")

	if goENV == "" || goENV == "development" {
		errThis := godotenv.Load()
		if errThis != nil {
			return errThis
		}
	}
	return nil
}
