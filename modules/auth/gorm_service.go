package auth

import (
	"base-site-api/models"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/dgrijalva/jwt-go"
	"github.com/gobuffalo/envy"
	"golang.org/x/crypto/bcrypt"
)

// GormService implementation Service interface with gorm.DB
type GormService struct {
	repository Repository
	signingKey []byte
}

// NewService return instance of GormService
func NewService(repository Repository) *GormService {
	signingKey, err := ioutil.ReadFile(envy.Get("JWT_KEY_PATH", ""))
	if err != nil {
		log.Fatal(err)
	}

	return &GormService{
		repository,
		signingKey,
	}
}

// Login func which return a new JWT token
func (s *GormService) Login(username string, password string) (string, error) {
	var tokenString string
	u, err := s.repository.FindUserByEmail(username)

	if err != nil {
		return tokenString, fmt.Errorf("could find user, %v", err)
	}

	pwdCompare := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	if pwdCompare != nil {
		return tokenString, fmt.Errorf("Error while comparing passwords %v", pwdCompare)
	}

	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Add(oneWeek()).Unix(),
		Issuer:    fmt.Sprintf("%s.api.go-with-jwt.it", envy.Get("GO_ENV", "development")),
		Id:        string(u.ID),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(s.signingKey)

	if err != nil {
		return tokenString, fmt.Errorf("could not sign token, %v", err)
	}

	return tokenString, nil
}

// UserInfo return necessary userinfo
func (s *GormService) UserInfo(userID uint) (*models.User, error) {
	return s.repository.Find(userID)
}

// ForgotPassword send email with ForgotPasswordToken
func (s *GormService) ForgotPassword(email string) error {
	return nil
}

// ResetPassword set new password for use if have valid token
func (s *GormService) ResetPassword(token string, newPassword string) error {
	return nil
}

// RegisterUser prepare, validate new user and save it to database
func (s *GormService) RegisterUser(u *models.User) error {
	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Passwords are not same")
	}

	ph, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("Error while hasing Passwords with error: %s", err)
	}

	u.PasswordHash = string(ph)
	u.Email = strings.ToLower(strings.TrimSpace(u.Email))

	return s.repository.CreateUser(u)
}

func oneWeek() time.Duration {
	return 7 * 24 * time.Hour
}
