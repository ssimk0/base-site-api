package announcement

import (
	"base-site-api/internal/app/models"
	"base-site-api/internal/log"
	"base-site-api/internal/tests/test_helper"
	"testing"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AnnouncementTestSuite struct {
	test_helper.RepositoryTestSuite
}

func (s *AnnouncementTestSuite) SetupTest() {
	s.Setup()
	s.Conn.Debug().AutoMigrate(
		&models.Announcement{},
	)
}

func (s *AnnouncementTestSuite) getTestAnnouncements() *models.Announcement {
	expire := (time.Now()).Add(time.Hour)
	return &models.Announcement{
		Message:  "Test",
		ExpireAt: expire,
	}
}

func (s *AnnouncementTestSuite) prepareTestData() []*models.Announcement {
	expire := (time.Now()).Add(time.Hour)
	announcements := []*models.Announcement{
		{
			Message:  "Test",
			ExpireAt: expire,
		},
		{
			Message:  "Test New",
			ExpireAt: expire,
		},
	}

	for _, a := range announcements {
		if err := s.Conn.Create(a).Error; err != nil {
			panic(err)
		}
	}

	return announcements
}

func (s *AnnouncementTestSuite) TestStore() {
	a := s.getTestAnnouncements()
	r := NewRepository(s.Conn)

	id, err := r.Store(a)

	if err != nil {
		s.T().Errorf("Error store annoucment %s", err)
	}
	log.Printf("ID store %d", id)
	assert.NotEqual(s.T(), 0, id)
}

func (s *AnnouncementTestSuite) TestActive() {
	s.prepareTestData()
	r := NewRepository(s.Conn)

	a, err := r.GetActive()

	if err != nil {
		s.T().Errorf("Error get active annoucment %s", err)
	}

	log.Debug(a)
	assert.Equal(s.T(), a.Message, "Test New")
}

func (s *AnnouncementTestSuite) TestActiveNoData() {
	r := NewRepository(s.Conn)

	a, err := r.GetActive()

	if err != nil {
		s.T().Errorf("Error get active annoucment %s", err)
	}

	assert.Equal(s.T(), a.Message, "")
	assert.Equal(s.T(), a.ID, uint(0))
}

func TestAnnouncementRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AnnouncementTestSuite))
}
