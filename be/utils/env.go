package utils

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func InitEnv() error {
	if envErr := godotenv.Load(); envErr != nil {
		return errors.New("Error loading .env file: " + envErr.Error())
	}

	return nil
}

func EnvString(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", errors.New("Unknown ENV variable: " + key)
	}

	return val, nil
}
func EnvInt(key string) (int, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return 0, errors.New("Unknown ENV variable: " + key)
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return 0, errors.New("Failed to parse ENV var: " + key + " = " + val)
	}

	return valInt, nil
}
