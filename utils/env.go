package utils

import (
	"os"
	"strconv"
)

func EnvString(key string, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}
func EnvInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valInt
}
