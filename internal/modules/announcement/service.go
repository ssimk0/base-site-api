package announcement

import "base-site-api/internal/app/models"

type service struct {
	repository Repository
}

// NewService return instance of service for article
func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetActive() (*models.Announcement, error) {
	return s.repository.GetActive()
}

func (s *service) Store(a *Announcement) (uint, error) {
	am := models.Announcement{
		Message:  a.Message,
		ExpireAt: a.ExpireAt,
	}

	return s.repository.Store(&am)
}
