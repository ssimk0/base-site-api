package modules

import (
	"base-site-api/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	Conn        *gorm.DB
	cleanupHook func()
}

func (s *RepositoryTestSuite) Setup() {
	var err error
	s.Conn, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		log.Fatal(err)
	}

	s.Conn.LogMode(true)
}

func (s *RepositoryTestSuite) BeforeTest(suiteName, testName string) {
	log.Debugf("Before test %s from suite %s", suiteName, testName)
	s.cleanupHook = utils.DeleteCreatedEntities(s.Conn)
}

func (s *RepositoryTestSuite) AfterTest(suiteName, testName string) {
	log.Debugf("After test %s from suite %s", suiteName, testName)
	s.cleanupHook()
}
