package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// ForgotPasswordToken model will used for auth
type ForgotPasswordToken struct {
	gorm.Model
	Token    string    `json:"token"`
	UserID   uint      `json:"-"`
	User     User      `json:"-"`
	ExpireAt time.Time `json:"expire_at"`
}
