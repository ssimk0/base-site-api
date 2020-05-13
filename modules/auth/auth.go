package auth

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// New setup whole module with all routes
func New(db *gorm.DB, api *fiber.Group) {
	handler := NewHandler(NewService(NewRepository(db)))

	api.Post("/v1/auth/login", handler.Login)
	api.Post("/v1/auth/register-user", handler.RegisterUser)
	// api.Get("/v1/auth/user", handler.GetUserInfo)
	api.Post("/v1/auth/forgot-password", handler.ForgotPassword)
	// api.Post("/v1/auth/reset-password", handler.ResetPassword)
}
