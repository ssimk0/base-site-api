package auth

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func New(db *gorm.DB, api *fiber.Group) {
	handler := NewHandler(NewService(NewRepository(db)))

	api.Post("/v1/auth/login", handler.Login)
	api.Get("/v1/auth/user", handler.GetUserInfo)
	api.Post("/v1/auth/forgot-password", handler.ForgotPassword)
	api.Post("/v1/auth/reset-password", handler.ResetPassword)
}
