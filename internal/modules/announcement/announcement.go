package announcement

import (
	"base-site-api/internal/models"
)

// Repository announcement
type Repository interface {
	GetActive() (*models.Announcement, error)
	Store(a *models.Announcement) (uint, error)
}
