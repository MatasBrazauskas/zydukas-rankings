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

	MaxBodyLength    = 1024 * 1024
	UserContextTitle = "userContext"
)

type CustomErrors struct {
	Error string `json:"error"`
}

var (
	ErrIntervalError  = CustomErrors{"Interval Error"}
	ErrReadingPayload = CustomErrors{"Error Reading Payload"}
	ErrWritingPayload = CustomErrors{"Error Writing Payload"}

	ErrCreatingUser = CustomErrors{"User Creation Error"}

	ErrNoAuthHeader = CustomErrors{"No Auth Header"}
	ErrInvalidJwt   = CustomErrors{"Invalid JWT token."}

	ErrNoFileAttached = CustomErrors{"No File attached."}
)
