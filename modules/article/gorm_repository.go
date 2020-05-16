package article

import (
	"base-site-api/models"
	// need to by for database
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

// GormRepository implementation of repository with gorm.DB
type GormRepository struct {
	db *gorm.DB
}

// NewRepository return instance of GormRepository
func NewRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}


// Find published article by slug
func (r *GormRepository) Find(id uint) (*models.Article, error) {
	article := models.Article{}

	if err := r.db.Set("gorm:auto_preload", true).Where("published = 1").First(&article, id).Error; err != nil {
		return nil, err
	}

	return &article, nil
}


// FindAll list all published articles also with order
func (r *GormRepository) FindAll(order string) ([]*models.Article, error) {
	var articles []*models.Article

	if err := r.db.Set("gorm:auto_preload", true).Model(&models.Article{}).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}


// Update the article
func (r *GormRepository) Update(article *models.Article, id uint) error {
	return r.db.Update(models.Article{
		Title:     article.Title,
		Body:      article.Body,
		Short:     article.Short,
		Slug:      article.Slug,
		Published: article.Published,
		Viewed:    article.Viewed,
	}).Error
}


// Store new article in db and return ID 
func (r *GormRepository) Store(article *models.Article, userID uint) (uint, error) {
	article.UserID = userID
	if err := r.db.Create(article).Error; err != nil {
		return 0, err
	}

	return article.ID, nil
}


// Delete article by ID 
func (r *GormRepository) Delete(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Article{}).Error; err != nil {
		return err
	}

	return nil
}
