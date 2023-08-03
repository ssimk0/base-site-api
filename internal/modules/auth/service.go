package auth

import (
	"base-site-api/internal/app/models"
	"base-site-api/internal/email"
	"base-site-api/internal/log"
	"base-site-api/internal/random"
	"bytes"
	"fmt"
	"html/template"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// service implementation ServiceI interface with gorm.DB
type service struct {
	repository Repository
	signingKey []byte
	templates  *template.Template
	emailer    email.Emailer
}

// NewService return instance of service
func NewService(repository Repository, signingKey []byte, templatePath string) Service {
	tpl := template.Must(template.ParseGlob(fmt.Sprintf("%s/emails/*.html", templatePath)))
	return &service{
		repository,
		signingKey,
		tpl,
		email.Instance(),
	}
}

// Login func which return a new JWT token
func (s *service) Login(username string, password string) (string, error) {
	var tokenString string
	u, err := s.repository.FindByEmail(username)

	if err != nil {
		return tokenString, fmt.Errorf("could find user, %v", err)
	}

	pwdCompare := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))

	if pwdCompare != nil {
		return tokenString, fmt.Errorf("error while comparing passwords %v", pwdCompare)
	}

	claims := jwt.MapClaims{
		"expireAt": time.Now().Add(oneWeek()).Unix(),
		"issuer":   "go-with-jwt.it",
		"id":       strconv.FormatUint(uint64(u.ID), 10),
		"subject":  strconv.FormatUint(uint64(u.ID), 10),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = token.SignedString(s.signingKey)

	if err != nil {
		return tokenString, fmt.Errorf("could not sign token, %v", err)
	}

	return tokenString, nil
}

// UserInfo return necessary userinfo
func (s *service) UserInfo(userID uint) (*models.User, error) {
	return s.repository.Find(userID)
}

type forgotPasswordMailData struct {
	FirstName string
	Link      string
}

// ForgotPassword send email with ForgotPasswordToken
func (s *service) ForgotPassword(email string, appUrl string) error {
	user, err := s.repository.FindByEmail(email)

	if err != nil {
		return err
	}

	token := &models.ForgotPasswordToken{
		Token:    random.String(10),
		UserID:   user.ID,
		ExpireAt: time.Now().Add(time.Minute * 15),
	}

	_, err = s.repository.StoreForgotPasswordToken(token)

	if err != nil {
		return err
	}

	w := &bytes.Buffer{}
	err = s.templates.ExecuteTemplate(w, "forgot-password.html", forgotPasswordMailData{
		FirstName: user.FirstName,
		Link:      fmt.Sprintf("%s/reset-password?token=%s", appUrl, token.Token),
	})

	if err != nil {
		return err
	}

	err = s.emailer.SendMail(email, w.Bytes())
	return err
}

// ResetPassword set new password for use if have valid token
func (s *service) ResetPassword(token string, newPassword string) error {
	t, err := s.repository.GetForgotPasswordToken(token)
	if err != nil {
		return err
	}

	if token == t.Token {

		u := &t.User
		pass, err := hashPassword(newPassword)
		if err != nil {
			return err
		}

		u.PasswordHash = pass
		err = s.repository.Update(u, u.ID)

		if err != nil {
			return err
		}
	}

	return nil
}

// RegisterUser prepare, validate new user and save it to database
func (s *service) RegisterUser(u *UserRequest) error {
	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("passwords are not same")
	}
	log.Errorf("%s %s %s", u.Email, u.FirstName, u.LastName)
	pass, err := hashPassword(u.Password)
	if err != nil {
		return err
	}
	user := models.User{
		PasswordHash: pass,
		Email:        strings.ToLower(strings.TrimSpace(u.Email)),
		FirstName:    u.FirstName,
		LastName:     u.LastName,
	}
	log.Errorf("%s %s %s", user.Email, user.FirstName, user.LastName)
	return s.repository.Store(&user)
}

func oneWeek() time.Duration {
	return 7 * 24 * time.Hour
}

func hashPassword(password string) (string, error) {
	ph, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error while hasing Passwords with error: %s", err)
	}

	return string(ph), nil
}
