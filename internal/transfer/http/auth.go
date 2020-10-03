package http

import (
	"base-site-api/internal/auth"
	"base-site-api/internal/common"
	"base-site-api/internal/dto"
	"base-site-api/internal/log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// AuthHandler auth
type AuthHandler struct {
	common.Handler
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
	r := &auth.LoginRequest{}

	if err := c.BodyParser(r); err != nil {
		log.Debugf("Wrong request login: %s", err)
		return h.Error(c, http.StatusForbidden)
	}

	token, err := h.service.Login(r.Username, r.Password)

	if err != nil {
		log.Debugf("Error while login in: %s", err)
		return h.Error(c, http.StatusForbidden)
	}

	return h.JSON(c, 200, &auth.TokenResponse{
		Token: token,
	})
}

// RegisterUser validate and register the user
func (h *AuthHandler) RegisterUser(c *fiber.Ctx) error {
	u := &auth.UserRequest{}

	if err := c.BodyParser(u); err != nil {
		return h.Error(c, 400)
	}

	err := h.service.RegisterUser(u)

	if err != nil {
		log.Errorf("Error while register user: %s", err)
		return h.Error(c, 400)
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

		return h.Error(c, 500)
	}

	return h.JSON(c, 200, auth.UserInfoResponse{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		IsAdmin:   u.IsAdmin,
		CanEdit:   u.CanEdit,
	})
}

// ForgotPassword based on email
func (h *AuthHandler) ForgotPassword(c *fiber.Ctx) error {
	u := &auth.User{}

	if err := c.BodyParser(u); err != nil {
		log.Errorf("Error while parsing body: %s", err)
		return h.Error(c, 400)
	}

	err := h.service.ForgotPassword(u.Email)

	if err != nil {
		log.Errorf("Error while processing forgot password: %s", err)

		return h.Error(c, 400)
	}

	return h.JSON(c, 200, &dto.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

// ResetPassword based on token from ForgotPassword
func (h *AuthHandler) ResetPassword(c *fiber.Ctx) error {
	u := &auth.ResetPasswordRequest{}
	token := c.Params("token")

	if err := c.BodyParser(u); err != nil {
		log.Errorf("Error while parsing reset password body: %s", err)
		return h.Error(c, 400)
	}

	if u.Password != u.PasswordConfirm {
		return h.Error(c, 400)
	}

	err := h.service.ResetPassword(token, u.Password)

	if err != nil {
		log.Errorf("Error while processing reset password: %s", err)

		return h.Error(c, 400)
	}

	return h.JSON(c, 200, dto.SuccessResponse{
		Success: true,
		ID:      0,
	})
}
