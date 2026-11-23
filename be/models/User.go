package models

import "time"

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique;not null"`
	FirebaseId string `gorm:"unique;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CreateUserPayload struct {
	Username string `json:"username" validate:"required,non_empty"`
}

type CreateUserContextPayload struct {
	Username   string
	FirebaseId string
}

type UserInfo struct {
	Username   string    `json:"username" validate:"required,non_empty"`
	FirebaseId string    `json:"firebase_id" validate:"required,non_empty"`
	CreatedAt  time.Time `json:"created_at" validate:"required,non_empty"`
	UpdatedAt  time.Time `json:"updated_at" validate:"required,non_empty"`
}

type UserContext struct {
	FireBaseId string `json:"firebase_id"`
}
