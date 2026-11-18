package utils

const (
	/*Port        = ":8080"
	DbAddress   = "postgres://admin:admin@localhost:5432/zyd-ranks?sslmode=disable"
	MaxOpenCons = 30
	MaxIdleCons = 30
	MaxIdleTime = "15m"

	WriteTimeout = 30 * time.Second
	ReadTimeout  = 30 * time.Second
	IdleTimeout  = 60 * time.Second*/

	UserContextTitle = "userContext"
)

type CustomErrors struct {
	Error string `json:"error"`
}

var (
	ErrCreatingUser       = CustomErrors{"User Creation Error"}
	ErrBadUserCredentials = CustomErrors{"Bad User Credentials"}
	ErrInvalidUserData    = CustomErrors{"Bad User Data"}
	ErrUserExists         = CustomErrors{"User already exists"}

	ErrNoAuthHeader = CustomErrors{"No Auth Header"}
	ErrInvalidJwt   = CustomErrors{"Invalid JWT token."}

	ErrNoFileAttached = CustomErrors{"No File attached."}
)
