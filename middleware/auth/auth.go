package auth

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

type Config struct {
	SigningKey []byte
	Filter func(c *fiber.Ctx) bool
}

func FilterOutGet(c *fiber.Ctx) bool {
	return c.Method() == "GET"
}

func New(cfg *Config) func(*fiber.Ctx) {
	// Return middleware handler
	return func(c *fiber.Ctx) {

		if cfg.Filter != nil && cfg.Filter(c) {
			c.Next()
			return
		}
	
		// Get authorization header
		tokenString := c.Get(fiber.HeaderAuthorization)

		if len(tokenString) == 0 {
			c.Status(http.StatusUnauthorized).Send(fmt.Errorf("No token set in headers"))
			return
		}

		if strings.HasPrefix(tokenString, "Bearer") {
			tokenString = tokenString[7:]
		}

		// parsing token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return cfg.SigningKey, nil
		})

		if err != nil {
			c.Status(http.StatusUnauthorized).Send(fmt.Errorf("Could not parse the token, %v", err))
			return
		}

		// getting claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Debug(claims["jti"])
			c.Locals("userID", claims["jti"].(string))
		} else {
			c.Status(http.StatusUnauthorized).Send(fmt.Errorf("Failed to validate token: %v", claims))
			return
		}
		c.Next()
	}
}
