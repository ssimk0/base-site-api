package handlers

import (
	"base-site-api/internal/app/dto"
	"base-site-api/internal/log"
	"base-site-api/internal/modules/announcement"
	"github.com/gofiber/fiber/v2"
)

type AnnouncementHandler struct {
	Handler
	service announcement.Service
}

func NewAnnouncementHandler(s announcement.Service) *AnnouncementHandler {
	return &AnnouncementHandler{
		service: s,
	}
}

func (h *AnnouncementHandler) Active(c *fiber.Ctx) error {
	a, err := h.service.GetActive()

	if err != nil {
		log.Errorf("Error while getting active announcement: %s", err)
		return h.Error(500)
	}

	return h.JSON(c, 200, &a)
}

func (h *AnnouncementHandler) Create(c *fiber.Ctx) error {
	data := &announcement.Announcement{}

	err := c.BodyParser(data)

	if err != nil {
		log.Debugf("Error while parsing announcement %s", err)
		return h.Error(400)
	}

	ID, err := h.service.Store(data)

	if err != nil {
		log.Errorf("Error while creating announcement: %s", err)
		return h.Error(500)
	}

	r := dto.SuccessResponse{
		Success: true,
		ID:      ID,
	}

	return h.JSON(c, 201, &r)
}
