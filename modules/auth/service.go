package auth

import (
	"base-site-api/models"
)

// Service interface for Auth
type Service interface {
	Login(username string, password string) (string, error)
	UserInfo(userID uint) (*models.User, error)
	RegisterUser(u *UserRequest) error
	ForgotPassword(email string) error
	ResetPassword(token string, newPassword string) error
}
