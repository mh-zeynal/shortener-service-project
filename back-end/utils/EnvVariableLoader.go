package utils

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func loadFile() error {
	err := godotenv.Load(".env")
	return err
}

func GetVariable(varName string) string {
	if loadFile() != nil {
		errors.New("error loading .env file")
	}
	return os.Getenv(varName)
}
