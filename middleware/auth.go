package middleware

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/envy"
	"github.com/gofiber/fiber"
)

var signingKey []byte

func init() {
	var err error
	signingKey, err = ioutil.ReadFile(envy.Get("JWT_KEY_PATH", ""))
	if err != nil {
		log.Fatal(err)
	}
}

func New() func(*fiber.Ctx) {

	// Return middleware handler
	return func(c *fiber.Ctx) {
	
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

			return signingKey, nil
		})

		if err != nil {
			c.Status(http.StatusUnauthorized).Send(fmt.Errorf("Could not parse the token, %v", err))
			return
		}

		// getting claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			c.Set("userId", claims["jti"].(string))
		} else {
			c.Status(http.StatusUnauthorized).Send(fmt.Errorf("Failed to validate token: %v", claims))
			return
		}
		c.Next()
	}
}
