package announcement

import "base-site-api/internal/app/models"

// Repository announcement
type Repository interface {
	GetActive() (*models.Announcement, error)
	Store(a *models.Announcement) (uint, error)
}

// Service announcement
type Service interface {
	GetActive() (*models.Announcement, error)
	Store(a *Announcement) (uint, error)
}
