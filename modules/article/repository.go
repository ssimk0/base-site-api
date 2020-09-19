package article

import (
	"base-site-api/models"
	"github.com/jinzhu/gorm"
)

//Repository interface for Article model
type Repository interface {
	Find(id uint) (*models.Article, error)
	FindBySlug(slug string) (*models.Article, error)
	FindAll(order string, offset int, limit int) ([]*models.Article, int, error)
	Update(article *models.Article, id uint) error
	Store(article *models.Article, userID uint) (uint, error)
	Delete(id uint) error
}

// repository implementation of repository with gorm.DB
type repository struct {
	db *gorm.DB
}

// NewRepository return instance of repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

// Find published article by id
func (r *repository) Find(id uint) (*models.Article, error) {
	article := models.Article{}

	if err := r.db.Set("gorm:auto_preload", true).Where("published = 1").First(&article, id).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

// FindBySlug published article by slug
func (r *repository) FindBySlug(slug string) (*models.Article, error) {
	article := models.Article{}

	if err := r.db.Set("gorm:auto_preload", true).Where("published = 1").Where("slug = ?", slug).First(&article).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

// FindAll list all published articles also with order
func (r *repository) FindAll(order string, offset int, limit int) ([]*models.Article, int, error) {
	var articles []*models.Article
	var count int

	if err := r.db.Set("gorm:auto_preload", true).Model(&models.Article{}).Order(order).Where("published = 1").Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Set("gorm:auto_preload", true).Model(&models.Article{}).Offset(offset).Limit(limit).Order(order).Where("published = 1").Find(&articles).Error; err != nil {
		return nil, 0, err
	}

	return articles, count, nil
}

// Update the article
func (r *repository) Update(article *models.Article, id uint) error {
	a, err := r.Find(id)
	if err != nil {
		return err
	}
	if article.Title != "" {
		a.Title = article.Title
	}
	if article.Body != "" {
		a.Body = article.Body
	}
	if article.Short != "" {
		a.Short = article.Short
	}
	if article.Slug != "" {
		a.Slug = article.Slug
	}
	if article.Viewed != 0 {
		a.Viewed = article.Viewed
	}
	a.Published = article.Published

	return r.db.Save(a).Error
}

// Store new article in db and return ID
func (r *repository) Store(article *models.Article, userID uint) (uint, error) {
	article.UserID = userID
	if err := r.db.Create(article).Error; err != nil {
		return 0, err
	}

	return article.ID, nil
}

// Delete article by ID
func (r *repository) Delete(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Article{}).Error; err != nil {
		return err
	}

	return nil
}
