package service

import (
	"base-site-api/models"
	"base-site-api/modules/article/repository"

	"github.com/prometheus/common/log"
)

// GormService implementation of article Service interface
type GormService struct {
	repository repository.Repository
}

func New(r repository.Repository) *GormService {
	return &GormService{
		repository: r,
	}
}

// Find return article by ID and increase viewed by 1
func (s *GormService) Find(id uint) (*models.Article, error) {
	article, err := s.repository.Find(id)

	if err != nil {
		return nil, err
	}

	(*article).Viewed += 1

	// update viewed is not critical error can be ignored
	_ = s.repository.Update(article, id)

	return article, nil
}

// FindAll articles and sort them by created_at or viewed
func (s *GormService) FindAll(sort string) ([]*models.Article, error) {
	order := "created_at desc"

	if sort == "top" {
		order = "viewed desc"
	}

	return s.repository.FindAll(order)
}

// Update simple update article
func (s *GormService) Update(article *models.Article, id uint) error {
	return s.repository.Update(article, id)
}

// Storre create a new article and return instance of it
func (s *GormService) Store(article *models.Article) (*models.Article, error) {
	id, err := s.repository.Store(article)

	if err != nil {
		return nil, err
	}

	return s.repository.Find(id)
}

// Delete article set the deleted_at and make it unavailable to retrieve
func (s *GormService) Delete(id uint, userId uint) error {
	log.Infof("Article with id %d deleted by user with id %d", id, userId)

	return s.repository.Delete(id)
}
