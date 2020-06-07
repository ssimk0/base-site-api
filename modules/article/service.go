package article

import (
	"base-site-api/models"
	"base-site-api/modules"
	"fmt"
	"github.com/gosimple/slug"

	log "github.com/sirupsen/logrus"
)

// Service interface for Article model
type Service interface {
	Find(slug string) (*models.Article, error)
	FindAll(sort string, page int, size int) ([]*models.Article, int, error)
	Update(article *models.Article, id uint) error
	Store(article *models.Article, userID uint) (*models.Article, error)
	Delete(id uint, userID uint) error
}

// service implementation of article Service interface
type service struct {
	modules.Service
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// Find return article by ID and increase viewed by 1
func (s *service) Find(slug string) (*models.Article, error) {
	article, err := s.repository.FindBySlug(slug)

	if err != nil {
		return nil, err
	}

	(*article).Viewed += 1

	// update viewed is not critical error can be ignored
	_ = s.repository.Update(article, article.ID)

	return article, nil
}

// FindAll articles and sort them by created_at or viewed
func (s *service) FindAll(sort string, page int, size int) ([]*models.Article, int, error) {
	order := "created_at desc"

	if sort == "viewed" || sort == "created_at" {
		order = fmt.Sprintf("%s desc", sort)
	}
	l, o := s.CalculateLimitAndOffset(page, size)
	return s.repository.FindAll(order, l, o)
}

// Update simple update article
func (s *service) Update(article *models.Article, id uint) error {
	return s.repository.Update(article, id)
}

// Store create a new article and return instance of it
func (s *service) Store(article *models.Article, userID uint) (*models.Article, error) {
	article.Slug = slug.Make(article.Title)
	id, err := s.repository.Store(article, userID)

	if err != nil {
		return nil, err
	}

	return s.repository.Find(id)
}

// Delete article set the deleted_at and make it unavailable to retrieve
// TODO: check if you have permission to delete
func (s *service) Delete(id uint, userID uint) error {
	log.Infof("Article with id %d deleted by user with id %d", id, userID)

	return s.repository.Delete(id)
}
