package auth

import (
	"base-site-api/models"
)

//Repository interface for Auth
type Repository interface {
	FindUserByEmail(email string) (*models.User, error)
	Find(id uint) (*models.User, error)
	Update(user *models.User, id uint) error
	StoreForgotPasswordToken(token string) (uint, error)
	GetForgotPasswordToken(token string) error // TODO: create model for token
}
