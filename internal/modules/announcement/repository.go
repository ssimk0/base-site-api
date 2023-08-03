package announcement

import (
	"base-site-api/internal/models"
	"github.com/jinzhu/gorm"
	"time"
)

type repository struct {
	db *gorm.DB
}

// NewRepository return instance of repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) GetActive() (*models.Announcement, error) {
	var a []*models.Announcement

	if err := r.db.Set("gorm:auto_preload", true).Model(&models.Announcement{}).Order("created_at desc").Where("expire_at > ?", time.Now()).Find(&a).Error; err != nil {
		return nil, err
	}

	if len(a) > 0 {
		return a[0], nil
	}

	return &models.Announcement{}, nil
}

func (r *repository) Store(a *models.Announcement) (uint, error) {

	if err := r.db.Create(a).Error; err != nil {
		return 0, err
	}

	return a.ID, nil
}
