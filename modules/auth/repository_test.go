package auth

import (
	"base-site-api/models"
	"base-site-api/modules"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AuthTestSuite struct {
	modules.RepositoryTestSuite
}

func (s *AuthTestSuite) SetupTest() {
	s.Setup()
	s.Conn.Debug().AutoMigrate(
		&models.User{},
		&models.ForgotPasswordToken{},
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

	user, err := r.FindUserByEmail(u.Email)

	if err != nil {
		s.T().Errorf("Error find user by email %s", err)
	}

	assert.Equal(s.T(), u.FirstName, user.FirstName)
	assert.Equal(s.T(), u.Email, user.Email)
}

func (s *AuthTestSuite) TestFindUserByWrongEmail() {
	r := NewRepository(s.Conn)

	user, _ := r.FindUserByEmail("example@example.com")

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
}

func (s *AuthTestSuite) TestCreateUser() {
	r := NewRepository(s.Conn)
	u := models.User{
		Email:        "test@test.com",
		FirstName:    "Hello",
		LastName:     "World",
		PasswordHash: "something",
	}

	err := r.StoreUser(&u)

	if err != nil {
		s.T().Errorf("Error store user %s", err)
	}

	user, err := r.FindUserByEmail(u.Email)

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

func (s *AuthTestSuite) TestStoreForgotToken() {
	r := NewRepository(s.Conn)
	t := models.ForgotPasswordToken{
		Token: "token",
	}

	tokenId, err := r.StoreForgotPasswordToken(&t)

	if err != nil {
		s.T().Errorf("Error find token %s", err)
	}
	token := &models.ForgotPasswordToken{}

	s.Conn.Find(token, tokenId)

	assert.Equal(s.T(), t.Token, token.Token)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
