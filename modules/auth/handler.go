package auth

import (
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"

	"github.com/gofiber/fiber"
	log "github.com/sirupsen/logrus"
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
func (h *Handler) Login(c *fiber.Ctx) {
	r := &LoginRequest{}

	if err := c.BodyParser(r); err != nil {
		log.Debugf("Wrong request login: %s", err)
		h.Error(c, 403)

		return
	}

	token, err := h.service.Login(r.Username, r.Password)

	if err != nil {
		log.Debugf("Error while login in: %s", err)
		h.Error(c, 403)

		return
	}

	h.JSON(c, 200, &TokenResponse{
		Token: token,
	})
}

// RegisterUser validate and register the user
func (h *Handler) RegisterUser(c *fiber.Ctx) {
	u := &UserRequest{}

	if err := c.BodyParser(u); err != nil {
		h.Error(c, 400)

		return
	}

	err := h.service.RegisterUser(u)

	if err != nil {
		log.Errorf("Error while register user: %s", err)
		h.Error(c, 400)

		return
	}

	h.JSON(c, 201, &responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

// GetUserInfo return all necessary information about the user based on JWT token
func (h *Handler) GetUserInfo(c *fiber.Ctx) {
	userID := h.ParseUserID(c)
	u, err := h.service.UserInfo(userID)
	if err != nil {
		log.Errorf("Error while getting the userinfo: %s", err)

		h.Error(c, 500)

		return
	}

	h.JSON(c, 200, UserInfoResponse{
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		IsAdmin:   u.IsAdmin,
		CanEdit:   u.CanEdit,
	})
}

// ForgotPassword based on email
func (h *Handler) ForgotPassword(c *fiber.Ctx) {
	u := &models.User{}

	if err := c.BodyParser(u); err != nil {
		log.Errorf("Error while parsing body: %s", err)
		h.Error(c, 400)

		return
	}

	err := h.service.ForgotPassword(u.Email)

	if err != nil {
		log.Errorf("Error while processing forgot password: %s", err)

		h.Error(c, 400)

		return

	}

	h.JSON(c, 200, &responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

// ResetPassword based on token from ForgotPassword
func (h *Handler) ResetPassword(c *fiber.Ctx) {
	u := &ResetPasswordRequest{}
	token := c.Params("token")

	if err := c.BodyParser(u); err != nil {
		log.Errorf("Error while parsing reset password body: %s", err)
		h.Error(c, 400)

		return
	}

	if u.Password != u.PasswordConfirm {
		h.Error(c, 400)

		return
	}

	err := h.service.ResetPassword(token, u.Password)

	if err != nil {
		log.Errorf("Error while processing reset password: %s", err)

		h.Error(c, 400)

		return
	}

	h.JSON(c, 200, responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}
