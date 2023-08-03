package handlers

import (
	"base-site-api/internal/app/dto"
	"base-site-api/internal/log"
	"base-site-api/internal/models"
	"base-site-api/internal/modules/auth"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// AuthHandler auth
type AuthHandler struct {
	Handler
	service auth.Service
}

// NewHandler return instance of AuthHandler
func NewAuthHandler(s auth.Service) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

// Login handler return the JWT token
func (h *AuthHandler) Login(c *fiber.Ctx) error {
	r := &models.LoginRequest{}

	if err := c.BodyParser(r); err != nil {
		log.Debugf("Wrong request login: %s", err)
		return h.Error(http.StatusForbidden)
	}

	token, err := h.service.Login(r.Username, r.Password)

	if err != nil {
		log.Debugf("Error while login in: %s", err)
		return h.Error(http.StatusForbidden)
	}

	return h.JSON(c, 200, &models.TokenResponse{
		Token: token,
	})
}

// RegisterUser validate and register the user
func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	u := &models.UserRequest{}

	if err := c.BodyParser(u); err != nil {
		return h.Error(400)
	}

	err := h.service.RegisterUser(u)

	if err != nil {
		log.Errorf("Error while register user: %s", err)
		return h.Error(400)
	}

	return h.JSON(c, 201, &dto.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

// GetUserInfo return all necessary information about the user based on JWT token
func (h *AuthHandler) GetUserInfo(c *fiber.Ctx) error {
	userID := h.ParseUserID(c)
	u, err := h.service.UserInfo(userID)
	if err != nil {
		log.Errorf("Error while getting the userinfo: %s", err)

		return h.Error(500)
	}

	return h.JSON(c, 200, models.UserInfoResponse{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		IsAdmin:   u.IsAdmin,
		CanEdit:   u.CanEdit,
	})
}

// ForgotPassword based on email
func (h *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	u := &models.User{}

	if err := c.BodyParser(u); err != nil {
		log.Errorf("Error while parsing body: %s", err)
		return h.Error(400)
	}

	err := h.service.ForgotPassword(u.Email, c.Locals("APP_URL").(string))

	if err != nil {
		log.Errorf("Error while processing forgot password: %s", err)

		return h.Error(400)
	}

	return h.JSON(c, 200, &dto.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

// ResetPassword based on token from ForgotPassword
func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	u := &models.ResetPasswordRequest{}
	token := c.Params("token")

	if err := c.BodyParser(u); err != nil {
		log.Errorf("Error while parsing reset password body: %s", err)
		return h.Error(400)
	}

	if u.Password != u.PasswordConfirm {
		return h.Error(400)
	}

	err := h.service.ResetPassword(token, u.Password)

	if err != nil {
		log.Errorf("Error while processing reset password: %s", err)

		return h.Error(400)
	}

	return h.JSON(c, 200, dto.SuccessResponse{
		Success: true,
		ID:      0,
	})
}
