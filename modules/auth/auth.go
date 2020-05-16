package auth

import (
	"base-site-api/config"
	"github.com/gofiber/fiber"
)

// New setup whole module with all routes
func New(config *config.Config, api *fiber.Group) {
	handler := NewHandler(NewService(NewRepository(config.Database), config.SigningKey))

	api.Post("/v1/auth/login", handler.Login)
	api.Post("/v1/auth/register-user", handler.RegisterUser)
	// api.Get("/v1/auth/user", handler.GetUserInfo)
	api.Post("/v1/auth/forgot-password", handler.ForgotPassword)
	api.Post("/v1/auth/reset-password/:token", handler.ResetPassword)
}
