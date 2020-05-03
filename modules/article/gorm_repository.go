package article

import (
	"base-site-api/models"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

func (r *GormRepository) Find(id uint) (*models.Article, error) {
	article := models.Article{}

	if err := r.db.Where("published = 1").First(&article, id).Error; err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *GormRepository) FindAll(order string) ([]*models.Article, error) {
	var articles []*models.Article

	if err := r.db.Model(&models.Article{}).Order(order).Where("published = 1").Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *GormRepository) Update(article *models.Article, id uint) error {
	err := r.db.Update(models.Article{
		Title:     article.Title,
		Body:      article.Body,
		Short:     article.Short,
		Slug:      article.Slug,
		Published: article.Published,
		Viewed:    article.Viewed,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *GormRepository) Store(article *models.Article) (uint, error) {
	if err := r.db.Create(article).Error; err != nil {
		return 0, err
	}

	return article.ID, nil
}

func (r *GormRepository) Delete(id uint) error {
	if err := r.db.Where("id = ?", id).Delete(&models.Article{}).Error; err != nil {
		return err
	}

	return nil
}
