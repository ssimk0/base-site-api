package auth

import (
	"base-site-api/models"
)

//Repository interface for Auth
type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	Find(id uint) (*models.User, error)
	Update(user *models.User, id uint) error
	CreateUser(user *models.User) error
	StoreForgotPasswordToken(token *models.ForgotPasswordToken) (uint, error)
	GetForgotPasswordToken(token string) (*models.ForgotPasswordToken, error)
}
