package auth

import (
	"base-site-api/models"
	"base-site-api/modules"
	"base-site-api/responses"

	"github.com/gofiber/fiber"
	"github.com/sirupsen/logrus"
)

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
		h.JSON(c, 403, &responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "AUTH_ERROR",
		})
		return
	}

	token, err := h.service.Login(r.Username, r.Password)

	if err != nil {
		h.JSON(c, 403, &responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "AUTH_ERROR",
		})
		return
	}

	h.JSON(c, 200, &TokenResponse{
		Token: token,
	})
}

func (h *Handler) RegisterUser(c *fiber.Ctx) {
	u := &UserRequest{}

	if err := c.BodyParser(u); err != nil {
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Bad parameters",
			Error:   "BAD_REQUEST",
		})
		return
	}

	err := h.service.RegisterUser(u)

	if err != nil {
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "BAD_REQUEST",
		})
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
		logrus.Errorf("Error while parsing body: %s", err)
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Bad parameters",
			Error:   "BAD_REQUEST",
		})
		return
	}

	err := h.service.ForgotPassword(u.Email)

	if err != nil {
		logrus.Errorf("Error while processing forgot password: %s", err)

		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Error while processing forgot password",
			Error:   "BAD_REQUEST",
		})
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
		logrus.Error(err)
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Bad parameters",
			Error:   "BAD_REQUEST",
		})
		return
	}

	if u.Password != u.PasswordConfirm {
		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Passwords are not same",
			Error:   "BAD_REQUEST",
		})
		return
	}

	err := h.service.ResetPassword(token, u.Password)

	if err != nil {
		logrus.Errorf("Error while processing forgot password: %s", err)

		h.JSON(c, 400, &responses.ErrorResponse{
			Message: "Error while processing forgot password",
			Error:   "BAD_REQUEST",
		})
		return
	}

	h.JSON(c, 200, responses.SuccessResponse{
		Success: true,
		ID:      0,
	})
}
