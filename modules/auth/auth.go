package auth

import (
	"base-site-api/config"
	"base-site-api/middleware/auth"
	"github.com/gofiber/fiber/v2"
)

func postOnlyFilter(c *fiber.Ctx) bool {
	return c.Method() == "POST"
}

type Auth struct {
}

// New setup whole module with all routes
func (m Auth) New(config *config.Config, api *fiber.Router) {
	handler := NewHandler(NewService(NewRepository(config.Database), config.SigningKey, config.Constants.TemplatePath))
	a := (*api).Group("/v1/auth")
	a.Use(auth.New(&auth.Config{
		SigningKey: config.SigningKey,
		Filter:     postOnlyFilter,
		DB:         config.Database,
	}))

	a.Get("/user", handler.GetUserInfo)
	a.Post("/login", handler.Login)
	a.Post("/register-user", handler.RegisterUser)
	a.Post("/forgot-password", handler.ForgotPassword)
	a.Post("/reset-password/:token", handler.ResetPassword)
}
