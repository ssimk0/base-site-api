package auth

import (
	"base-site-api/responses"
	"base-site-api/models"
	"github.com/sirupsen/logrus"
	"github.com/gofiber/fiber"
)

type AuthHandler struct {
	service Service
}

func NewHandler(s Service) *AuthHandler {
	return &AuthHandler{
		service: s,
	}
}

type LoginRequest struct {
	Username string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(c *fiber.Ctx) {
	r := LoginRequest{}

	if err := c.BodyParser(r); err != nil {
		c.Status(403).JSON(&responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "AUTH_ERROR",
		})
	}

	token, err := h.service.Login(r.Username, r.Password)

	if err != nil {
		c.Status(403).JSON(&responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "AUTH_ERROR",
		})
	}

	c.JSON(&TokenResponse{
		Token: token,
	})
}

func (h *AuthHandler) RegisterUser(c *fiber.Ctx) {
	u := models.User{}

	if err := c.BodyParser(u); err != nil {
		logrus.Error(err)
		c.Status(400).JSON(&responses.ErrorResponse{
			Message: "Bad parameters",
			Error:   "BAD_REQUEST",
		})
	}
	logrus.Debugf("%v", u)
	err := h.service.RegisterUser(&u)

	if err != nil {
		logrus.Error(err)
		c.Status(403).JSON(&responses.ErrorResponse{
			Message: "Wrong password or username",
			Error:   "BAD_REQUEST",
		})
	}

	c.Status(201).JSON(responses.SuccessResponse{
		Success: true,
		Id: 0,
	})
}