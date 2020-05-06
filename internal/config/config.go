package config

import (
	"os"

	"base-site-api/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Constants for whole application setup
type Constants struct {
	ADDRESS string
	ENV     string
}

// Config application with all constants and database
type Config struct {
	Constants
	Database *gorm.DB
}

func initDB(env string) (*gorm.DB, error) {
	conn, err := gorm.Open("mysql", os.Getenv("DB_URI"))
	if err != nil {
		return nil, err
	}

	conn.LogMode(env == "development")
	conn.Debug().AutoMigrate(
		&models.Article{},
		&models.User{},
		&models.ForgotPasswordToken{},
	)

	return conn, nil
}

// New return application Config
func New() (*Config, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	constants := Constants{}
	constants.ENV = os.Getenv("GO_ENV")
	constants.ADDRESS = os.Getenv("ADDRESS")

	db, err := initDB(constants.ENV)

	if err != nil {
		return nil, err
	}

	return &Config{
		Constants: constants,
		Database:  db,
	}, nil
}
