package models

import (
	"time"
)

// User model will used for auth
type User struct {
	DatabaseModel
	Email        string `json:"email" gorm:"type:varchar(255);unique_index;not null"`
	FirstName    string `json:"first_name" gorm:"not null"`
	LastName     string `json:"last_name" gorm:"not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	CanEdit      bool   `json:"-"`
	IsAdmin      bool   `json:"-"`
}

// ForgotPasswordToken model will used for auth
type ForgotPasswordToken struct {
	DatabaseModel
	Token    string    `json:"token"`
	UserID   uint      `json:"-"`
	User     User      `json:"-"`
	ExpireAt time.Time `json:"expire_at"`
}

// UserRequest struct handle params for register user
type UserRequest struct {
	Email           string `json:"email"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation"`
}

// ResetPasswordRequest handle params for reset password
type ResetPasswordRequest struct {
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirmation"`
}

// UserInfoResponse struct return all needed params
type UserInfoResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsAdmin   bool   `json:"is_admin"`
	CanEdit   bool   `json:"can_edit"`
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
