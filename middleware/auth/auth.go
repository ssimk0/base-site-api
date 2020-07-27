package auth

import (
	"base-site-api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
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
func New(cfg *Config) func(*fiber.Ctx) {
	// Return middleware handler
	return func(c *fiber.Ctx) {

		if cfg.Filter != nil && cfg.Filter(c) {
			c.Next()
			return
		}

		// Get authorization header
		tokenString := c.Get(fiber.HeaderAuthorization)
		if len(tokenString) < 7 {
			c.Status(http.StatusUnauthorized).Send(fmt.Errorf("no token set in headers"))
			return
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
			c.Status(http.StatusUnauthorized).Send(fmt.Errorf("could not parse the token, %v", err))
			return
		}

		// getting claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			userID, err := strconv.ParseUint(claims["jti"].(string), 10, 32)
			if err != nil {
				c.Status(http.StatusUnauthorized).Send(fmt.Errorf("failed to validate token: %v", claims))
				return
			}

			user := models.User{}

			if err := cfg.DB.First(&user, userID).Error; err != nil {
				c.Status(http.StatusUnauthorized).Send(fmt.Errorf("failed to validate token: %v", claims))
				return
			}

			if user.CanEdit || user.IsAdmin {
				c.Locals("userID", uint(userID))
			} else {

				c.Status(http.StatusUnauthorized).Send(fmt.Errorf("failed to validate token: %v", claims))
				return
			}

		} else {
			c.Status(http.StatusUnauthorized).Send(fmt.Errorf("failed to validate token: %v", claims))
			return
		}
		c.Next()
	}
}
