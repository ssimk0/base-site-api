package repository

import (
	"base-site-api/models"
)

//Repository interface for Article model
type Repository interface {
	Find(id uint) (*models.Article, error)
	FindAll(order string) ([]*models.Article, error)
	Update(article *models.Article, id uint) error
	Store(article *models.Article) (uint, error)
	Delete(id uint) error
}
