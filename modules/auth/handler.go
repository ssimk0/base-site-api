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

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

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

func (h *Handler) ResetPassword(c *fiber.Ctx) {
	u := UserRequest{}
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
		log.Errorf("Error while processing forgot password: %s", err)

		h.Error(c, 400)

		return
	}

	h.JSON(c, 200, responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}
