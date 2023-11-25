package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPWD(password string)(string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", fmt.Errorf("password hash generation failure: %v", err)
	}

	return string(bytes), nil
}
