package auth

import (
	"base-site-api/internal/models"
	"github.com/jinzhu/gorm"
)

// repository implementation Repository with gorm.DB
type repository struct {
	db *gorm.DB
}

// NewRepository return the new repository
func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

// FindByEmail return User find by email
func (r *repository) FindByEmail(email string) (*models.User, error) {
	user := models.User{}

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Find return User by id
func (r *repository) Find(id uint) (*models.User, error) {
	user := models.User{}

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update instance the user
func (r *repository) Update(user *models.User, id uint) error {
	u, err := r.Find(id)
	if err != nil {
		return err
	}

	if user.Email != "" {
		u.Email = user.Email
	}

	if user.PasswordHash != "" {
		u.PasswordHash = user.PasswordHash
	}

	if user.LastName != "" {
		u.LastName = user.LastName
	}

	if user.CanEdit {
		u.CanEdit = user.CanEdit
	}

	if user.IsAdmin {
		u.IsAdmin = user.IsAdmin
	}

	return r.db.Save(u).Error
}

// StoreForgotPasswordToken and return id
func (r *repository) StoreForgotPasswordToken(token *models.ForgotPasswordToken) (uint, error) {
	if err := r.db.Create(token).Error; err != nil {
		return 0, err
	}

	return token.ID, nil
}

// GetForgotPasswordToken by token
func (r *repository) GetForgotPasswordToken(token string) (*models.ForgotPasswordToken, error) {
	t := models.ForgotPasswordToken{}

	if err := r.db.Set("gorm:auto_preload", true).Where("token = ?", token).First(&t).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

// Store new instance
func (r *repository) Store(user *models.User) error {
	return r.db.Create(user).Error
}
