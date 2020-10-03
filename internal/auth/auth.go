package auth

import (
	"base-site-api/internal/common"
	"time"
)

// User model will used for auth
type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	DeletedAt    *time.Time `json:"-" sql:"index"`
	CreatedAt    time.Time  `json:"-"`
	UpdatedAt    time.Time  `json:"-"`
	Email        string     `json:"email" gorm:"type:varchar(255);unique_index;not null"`
	FirstName    string     `json:"first_name" gorm:"not null"`
	LastName     string     `json:"last_name" gorm:"not null"`
	PasswordHash string     `json:"-" gorm:"not null"`
	CanEdit      bool       `json:"-"`
	IsAdmin      bool       `json:"-"`
}

// ForgotPasswordToken model will used for auth
type ForgotPasswordToken struct {
	common.Model
	Token    string    `json:"token"`
	UserID   uint      `json:"-"`
	User     User      `json:"-"`
	ExpireAt time.Time `json:"expire_at"`
}

// Repository interface of auth
type Repository interface {
	FindUserByEmail(email string) (*User, error)
	Find(id uint) (*User, error)
	Update(user *User, id uint) error
	StoreUser(user *User) error
	StoreForgotPasswordToken(token *ForgotPasswordToken) (uint, error)
	GetForgotPasswordToken(token string) (*ForgotPasswordToken, error)
}

// Service interface for Auth
type Service interface {
	Login(username string, password string) (string, error)
	UserInfo(userID uint) (*User, error)
	RegisterUser(u *UserRequest) error
	ForgotPassword(email string) error
	ResetPassword(token string, newPassword string) error
}
