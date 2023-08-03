package auth

import (
	"base-site-api/internal/models"
	"base-site-api/internal/tests/test_helper"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	test_helper.RepositoryTestSuite
}

func (s *AuthTestSuite) SetupTest() {
	s.Setup()
	s.Conn.Debug().AutoMigrate(
		&models.ForgotPasswordToken{},
		&models.User{},
	)
}

func (s *AuthTestSuite) prepareTestData() (*models.User, *models.ForgotPasswordToken) {
	psw, err := hashPassword("password")

	if err != nil {
		panic(err)
	}

	user := &models.User{
		FirstName:    "test",
		LastName:     "user",
		Email:        "simko22@gmail.com",
		PasswordHash: psw,
	}
	if err := s.Conn.Create(user).Error; err != nil {
		panic(err)
	}

	token := &models.ForgotPasswordToken{
		Token:    psw,
		UserID:   user.ID,
		ExpireAt: time.Now().Add(5000),
	}

	if err := s.Conn.Create(token).Error; err != nil {
		panic(err)
	}

	return user, token
}

func (s *AuthTestSuite) TestFindUserByEmail() {
	u, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	user, err := r.FindByEmail(u.Email)

	if err != nil {
		s.T().Errorf("Error find user by email %s", err)
	}

	assert.Equal(s.T(), u.FirstName, user.FirstName)
	assert.Equal(s.T(), u.Email, user.Email)
}

func (s *AuthTestSuite) TestFindUserByWrongEmail() {
	r := NewRepository(s.Conn)

	user, _ := r.FindByEmail("example@example.com")

	assert.Nil(s.T(), user)
}

func (s *AuthTestSuite) TestFindUserByID() {
	u, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	user, err := r.Find(u.ID)

	if err != nil {
		s.T().Errorf("Error find user by id %s", err)
	}

	assert.Equal(s.T(), u.FirstName, user.FirstName)
	assert.Equal(s.T(), u.Email, user.Email)
}

func (s *AuthTestSuite) TestFindUserByWrongID() {
	r := NewRepository(s.Conn)

	user, _ := r.Find(99999)

	assert.Nil(s.T(), user)
}

func (s *AuthTestSuite) TestUpdateUser() {
	u, _ := s.prepareTestData()
	r := NewRepository(s.Conn)
	u.FirstName = "new first Name"
	u.LastName = "married"
	u.CanEdit = true
	u.IsAdmin = true

	err := r.Update(u, u.ID)

	if err != nil {
		s.T().Errorf("Error updating user %s", err)
	}

	user, err := r.Find(u.ID)

	if err != nil {
		s.T().Errorf("Error find user by id %s", err)
	}

	// FirstName is immutable
	assert.NotEqual(s.T(), u.FirstName, user.FirstName)
	assert.Equal(s.T(), "married", user.LastName)
	assert.Equal(s.T(), true, user.CanEdit)
	assert.Equal(s.T(), true, user.IsAdmin)
}

func (s *AuthTestSuite) TestUpdateUserNotFound() {
	u, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	err := r.Update(u, 0)

	assert.NotNil(s.T(), err)
}

func (s *AuthTestSuite) TestCreateUser() {
	r := NewRepository(s.Conn)
	u := models.User{
		Email:        "test@test.com",
		FirstName:    "Hello",
		LastName:     "World",
		PasswordHash: "something",
	}

	err := r.Store(&u)

	if err != nil {
		s.T().Errorf("Error store user %s", err)
	}

	user, err := r.FindByEmail(u.Email)

	if err != nil {
		s.T().Errorf("Error find user by email %s", err)
	}

	assert.Equal(s.T(), u.FirstName, user.FirstName)
	assert.Equal(s.T(), u.Email, user.Email)
	// this is default values
	assert.Equal(s.T(), false, user.IsAdmin)
	assert.Equal(s.T(), false, user.CanEdit)
}

func (s *AuthTestSuite) TestGetForgotToken() {
	r := NewRepository(s.Conn)
	_, t := s.prepareTestData()

	token, err := r.GetForgotPasswordToken(t.Token)

	if err != nil {
		s.T().Errorf("Error find token %s", err)
	}

	assert.Equal(s.T(), t.Token, token.Token)
	assert.Equal(s.T(), t.UserID, token.UserID)
	assert.Equal(s.T(), "simko22@gmail.com", token.User.Email)
}

func (s *AuthTestSuite) TestGetForgotTokenNotFound() {
	r := NewRepository(s.Conn)

	_, err := r.GetForgotPasswordToken("12234")

	assert.NotNil(s.T(), err)
}

func (s *AuthTestSuite) TestStoreForgotToken() {
	r := NewRepository(s.Conn)
	t := models.ForgotPasswordToken{
		Token: "token",
	}

	tokenID, err := r.StoreForgotPasswordToken(&t)

	if err != nil {
		s.T().Errorf("Error find token %s", err)
	}
	token := &models.ForgotPasswordToken{}

	s.Conn.Find(token, tokenID)

	assert.Equal(s.T(), t.Token, token.Token)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
