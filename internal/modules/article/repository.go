package article

import (
	"base-site-api/internal/log"
	"base-site-api/internal/models"
	"base-site-api/internal/pagination"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

// repository implementation of repository with gorm.DB
type repository struct {
	pagination.Service
	db *gorm.DB
}

// NewRepository return instance of repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		pagination.Service{},
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
func (r *repository) FindAll(sort string, page int, size int) ([]*models.Article, int, error) {
	order := "created_at desc"

	if sort == "viewed" || sort == "created_at" {
		order = fmt.Sprintf("%s desc", sort)
	}
	offset, limit := r.CalculateLimitAndOffset(page, size)

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
	if article.Image != "" {
		a.Image = article.Image
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
	article.Slug = slug.Make(article.Title)
	if err := r.db.Create(article).Error; err != nil {
		return 0, err
	}

	return article.ID, nil
}

// Delete article by ID
func (r *repository) Delete(id uint, userID uint) error {
	log.Infof("Article with id %d deleted by user with id %d", id, userID)
	if err := r.db.Where("id = ?", id).Delete(&models.Article{}).Error; err != nil {
		return err
	}

	return nil
}
