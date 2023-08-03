package middleware

import (
	"base-site-api/internal/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// Config for auth middleware
type Config struct {
	SigningKey []byte
	Filter     func(c *fiber.Ctx) bool
	DB         *gorm.DB
}

// FilterGetOnly filter out request based on method GET
func FilterGetOnly(c *fiber.Ctx) bool {
	return c.Method() == "GET"
}

// New return the middleware function
func NewAuthMiddleware(cfg *Config) fiber.Handler {
	// Return middleware handler
	return func(c *fiber.Ctx) error {

		if cfg.Filter != nil && cfg.Filter(c) {
			return c.Next()
		}

		// Get authorization header
		tokenString := c.Get(fiber.HeaderAuthorization)
		if len(tokenString) < 7 {
			return fiber.NewError(http.StatusUnauthorized, "no token set in headers")
		}

		if strings.HasPrefix(tokenString, "Bearer") {
			tokenString = tokenString[7:]
		}

		// parsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return cfg.SigningKey, nil
		})

		if err != nil {
			return fiber.NewError(http.StatusUnauthorized, fmt.Sprintf("could not parse the token, %v", err))

		}

		// getting claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, err := strconv.ParseUint(claims["jti"].(string), 10, 32)
			if err != nil {
				return fiber.NewError(http.StatusUnauthorized, fmt.Sprintf("failed to validate token: %v", claims))
			}

			user := models.User{}

			if err := cfg.DB.First(&user, userID).Error; err != nil {
				return fiber.NewError(http.StatusUnauthorized, fmt.Sprintf("failed to validate token: %v", claims))
			}

			if user.CanEdit || user.IsAdmin {
				c.Locals("userID", uint(userID))
			} else {

				return fiber.NewError(http.StatusUnauthorized, fmt.Sprintf("failed to validate token: %v", claims))
			}

		} else {
			return fiber.NewError(http.StatusUnauthorized, fmt.Sprintf("failed to validate token: %v", claims))
		}
		return c.Next()
	}
}
