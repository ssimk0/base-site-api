package auth

import (
	"base-site-api/config"
	"base-site-api/middleware/auth"
	"github.com/gofiber/fiber"
)

func postOnlyFilter(c *fiber.Ctx) bool {
	return c.Method() == "POST"
}

// New setup whole module with all routes
func New(config *config.Config, api *fiber.Group) {
	handler := NewHandler(NewService(NewRepository(config.Database), config.SigningKey))
	a := api.Group("/v1/auth")
	a.Use(auth.New(&auth.Config{
		SigningKey: config.SigningKey,
		Filter:     postOnlyFilter,
	}))

	a.Get("/user", handler.GetUserInfo)
	a.Post("/login", handler.Login)
	a.Post("/register-user", handler.RegisterUser)
	a.Post("/forgot-password", handler.ForgotPassword)
	a.Post("/reset-password/:token", handler.ResetPassword)
}
