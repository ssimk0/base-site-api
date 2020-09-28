package auth

import (
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"

	"base-site-api/log"

	"github.com/gofiber/fiber/v2"
)

// Handler auth
type Handler struct {
	modules.Handler
	service Service
}

// NewHandler return instance of Handler
func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

// Login handler return the JWT token
func (h *Handler) Login(c *fiber.Ctx) error {
	r := &LoginRequest{}

	if err := c.BodyParser(r); err != nil {
		log.Debugf("Wrong request login: %s", err)
		return h.Error(c, 403)
	}

	token, err := h.service.Login(r.Username, r.Password)

	if err != nil {
		log.Debugf("Error while login in: %s", err)
		return h.Error(c, 403)
	}

	return h.JSON(c, 200, &TokenResponse{
		Token: token,
	})
}

// RegisterUser validate and register the user
func (h *Handler) RegisterUser(c *fiber.Ctx) error {
	u := &UserRequest{}

	if err := c.BodyParser(u); err != nil {
		return h.Error(c, 400)
	}

	err := h.service.RegisterUser(u)

	if err != nil {
		log.Errorf("Error while register user: %s", err)
		return h.Error(c, 400)
	}

	return h.JSON(c, 201, &responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

// GetUserInfo return all necessary information about the user based on JWT token
func (h *Handler) GetUserInfo(c *fiber.Ctx) error {
	userID := h.ParseUserID(c)
	u, err := h.service.UserInfo(userID)
	if err != nil {
		log.Errorf("Error while getting the userinfo: %s", err)

		return h.Error(c, 500)
	}

	return h.JSON(c, 200, UserInfoResponse{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		IsAdmin:   u.IsAdmin,
		CanEdit:   u.CanEdit,
	})
}

// ForgotPassword based on email
func (h *Handler) ForgotPassword(c *fiber.Ctx) error {
	u := &models.User{}

	if err := c.BodyParser(u); err != nil {
		log.Errorf("Error while parsing body: %s", err)
		return h.Error(c, 400)
	}

	err := h.service.ForgotPassword(u.Email)

	if err != nil {
		log.Errorf("Error while processing forgot password: %s", err)

		return h.Error(c, 400)
	}

	return h.JSON(c, 200, &responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

// ResetPassword based on token from ForgotPassword
func (h *Handler) ResetPassword(c *fiber.Ctx) error {
	u := &ResetPasswordRequest{}
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

	return h.JSON(c, 200, responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}
