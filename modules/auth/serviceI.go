package auth

import (
	"base-site-api/models"
)

// ServiceI interface for Auth
type ServiceI interface {
	Login(username string, password string) (string, error)
	UserInfo(userID uint) (*models.User, error)
	RegisterUser(u *UserRequest) error
	ForgotPassword(email string) error
	ResetPassword(token string, newPassword string) error
}
