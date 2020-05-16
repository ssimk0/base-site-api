package auth

import (
	"base-site-api/models"

	"github.com/jinzhu/gorm"
)

// GormRepository implementation Repository with gorm.DB
type GormRepository struct {
	db *gorm.DB
}

// NewRepository return the new GormRepository
func NewRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{
		db,
	}
}

// FindUserByEmail return User find by email
func (r *GormRepository) FindUserByEmail(email string) (*models.User, error) {
	user := models.User{}

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Find return User by id
func (r *GormRepository) Find(id uint) (*models.User, error) {
	user := models.User{}

	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update instance the user
func (r *GormRepository) Update(user *models.User, id uint) error {

	return r.db.Update(models.User{
		ID:           id,
		Email:        user.Email,
		FirstName:    user.LastName,
		PasswordHash: user.PasswordHash,
		CanEdit:      user.CanEdit, // TODO: make sure that is only changed by admin
		IsAdmin:      user.IsAdmin,
	}).Error
}

// StoreForgotPasswordToken and return id
func (r *GormRepository) StoreForgotPasswordToken(token *models.ForgotPasswordToken) (uint, error) {
	if err := r.db.Create(token).Error; err != nil {
		return 0, err
	}

	return token.ID, nil
}

// GetForgotPasswordToken by token
func (r *GormRepository) GetForgotPasswordToken(token string) (*models.ForgotPasswordToken, error) {
	t := models.ForgotPasswordToken{}

	if err := r.db.Set("gorm:auto_preload", true).Where("token = ?", token).First(&t).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

// CreateUser new instance
func (r *GormRepository) CreateUser(user *models.User) error {
	return r.db.Create(user).Error
}
