package auth

// UserRequest struct handle params for register user
type UserRequest struct {
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation"`
}

// LoginRequest struct handle params for login
type LoginRequest struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

// TokenResponse only json token reponse
type TokenResponse struct {
	Token string `json:"token"`
}
