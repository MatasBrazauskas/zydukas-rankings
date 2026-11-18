package models

type User struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique;not null"`
	FirebaseId string `gorm:"unique;not null"`
}

type CreateUserPayload struct {
	Username string `json:"username" validate:"required,non_empty"`
}

type CreateUserContextPayload struct {
	Username   string
	FirebaseId string
}

type UserContext struct {
	FireBaseId string `json:"firebase_id"`
}
