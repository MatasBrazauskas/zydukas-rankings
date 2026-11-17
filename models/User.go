package models

type User struct {
	Username string `json:"username"`
}

type CreateUserPayload struct {
	Username string `json:"username" validate:"required,non_empty"`
}

type UserContext struct {
	Username string
	Email    string
}
