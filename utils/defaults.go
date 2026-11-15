package utils

import (
	"time"
)

const (
	Port        = ":8080"
	DbAddress   = "postgres://admin:admin@localhost:5432/zyd-ranks?sslmode=disable"
	MaxOpenCons = 30
	MaxIdleCons = 30
	MaxIdleTime = "15m"

	WriteTimeout = 30 * time.Second
	ReadTimeout  = 30 * time.Second
	IdleTimeout  = 60 * time.Second
)

type CustomErrors struct {
	Error string `json:"error"`
}

var (
	ErrIntervalError = CustomErrors{"interval Error"}

	ErrCreatingUser = CustomErrors{"user creation error"}
)
