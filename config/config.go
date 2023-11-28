package config

import (
	"os"

	"github.com/joho/godotenv"
)

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

func GetENV(key string) string {
	err1 := godotenv.Load()
	if err1 != nil {
		return err1.Error()
	}
	
	return os.Getenv(key)
}
