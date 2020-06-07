package auth

import (
	"base-site-api/models"
	"base-site-api/responses"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

type Handler struct {
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
		c.Status(403).JSON(&responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "AUTH_ERROR",
		})
		return
	}

	token, err := h.service.Login(r.Username, r.Password)

	if err != nil {
		c.Status(403).JSON(&responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "AUTH_ERROR",
		})
		return
	}

	c.JSON(&TokenResponse{
		Token: token,
	})
}

func (h *Handler) RegisterUser(c *fiber.Ctx) {
	u := &UserRequest{}

	if err := c.BodyParser(u); err != nil {
		c.Status(400).JSON(&responses.ErrorResponse{
			Message: "Bad parameters",
			Error:   "BAD_REQUEST",
		})
		return
	}

	err := h.service.RegisterUser(u)

	if err != nil {
		c.Status(403).JSON(&responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "BAD_REQUEST",
		})
		return
	}

	c.Status(201).JSON(responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

func (h *Handler) ForgotPassword(c *fiber.Ctx) {
	u := &models.User{}

	if err := c.BodyParser(u); err != nil {
		logrus.Errorf("Error while parsing body: %s", err)
		c.Status(400).JSON(&responses.ErrorResponse{
			Message: "Bad parameters",
			Error:   "BAD_REQUEST",
		})
		return
	}

	err := h.service.ForgotPassword(u.Email)

	if err != nil {
		logrus.Errorf("Error while processing forgot password: %s", err)

		c.Status(400).JSON(&responses.ErrorResponse{
			Message: "Error while processing forgot password",
			Error:   "BAD_REQUEST",
		})
		return

	}

	c.JSON(responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}

func (h *Handler) ResetPassword(c *fiber.Ctx) {
	u := UserRequest{}
	token := c.Params("token")

	if err := c.BodyParser(u); err != nil {
		logrus.Error(err)
		c.Status(400).JSON(&responses.ErrorResponse{
			Message: "Bad parameters",
			Error:   "BAD_REQUEST",
		})
		return
	}

	if u.Password != u.PasswordConfirm {
		c.Status(400).JSON(&responses.ErrorResponse{
			Message: "Passwords are not same",
			Error:   "BAD_REQUEST",
		})
		return
	}

	err := h.service.ResetPassword(token, u.Password)

	if err != nil {
		logrus.Errorf("Error while processing forgot password: %s", err)

		c.Status(400).JSON(&responses.ErrorResponse{
			Message: "Error while processing forgot password",
			Error:   "BAD_REQUEST",
		})
		return
	}

	c.JSON(responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}
