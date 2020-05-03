package auth

import (
	"base-site-api/models"
)

// Service interface for Auth
type Service interface {
	Login(username string, password string) (*models.User, error)
	UserInfo(userId uint) (*models.User, error)
	ForgotPassword(email string) error
	ResetPassword(token string, newPassword string) error
}
