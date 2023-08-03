package auth

import (
	"base-site-api/internal/models"
)

// Repository interface of auth
type Repository interface {
	FindByEmail(email string) (*models.User, error)
	Find(id uint) (*models.User, error)
	Update(user *models.User, id uint) error
	Store(user *models.User) error
	StoreForgotPasswordToken(token *models.ForgotPasswordToken) (uint, error)
	GetForgotPasswordToken(token string) (*models.ForgotPasswordToken, error)
}

// Service interface for Auth
type Service interface {
	Login(username string, password string) (string, error)
	UserInfo(userID uint) (*models.User, error)
	RegisterUser(u *models.UserRequest) error
	ForgotPassword(email string, appUrl string) error
	ResetPassword(token string, newPassword string) error
}
