package utils

const (
	UserContextTitle = "userContext"
)

type CustomErrors struct {
	Error string `json:"error"`
}

var (
	ErrCreatingUser          = CustomErrors{"User Creation Error"}
	ErrBadUserCredentials    = CustomErrors{"Bad User Credentials"}
	ErrInvalidUserData       = CustomErrors{"Bad User Data"}
	ErrUserExists            = CustomErrors{"User already exists"}
	ErrUserNotFound          = CustomErrors{"User not found"}
	ErrUpdatingUsersUsername = CustomErrors{"Error Updating Username"}

	ErrNoAuthHeader = CustomErrors{"No Auth Header"}
	ErrInvalidJwt   = CustomErrors{"Invalid JWT token."}

	ErrNoFileAttached = CustomErrors{"No File attached."}
	ErrCreatingPosts  = CustomErrors{"Error Creating Posts"}
)
