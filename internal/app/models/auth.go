package models

import (
	"time"
)

// User model will used for auth
type User struct {
	Model
	Email        string `json:"email" gorm:"type:varchar(255);unique_index;not null"`
	FirstName    string `json:"first_name" gorm:"not null"`
	LastName     string `json:"last_name" gorm:"not null"`
	PasswordHash string `json:"-" gorm:"not null"`
	CanEdit      bool   `json:"-"`
	IsAdmin      bool   `json:"-"`
}

// ForgotPasswordToken model will used for auth
type ForgotPasswordToken struct {
	Model
	Token    string    `json:"token"`
	UserID   uint      `json:"-"`
	User     User      `json:"-"`
	ExpireAt time.Time `json:"expire_at"`
}
