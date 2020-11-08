package test_helper

import (
	"base-site-api/internal/app/config"
	"base-site-api/internal/database"
	"base-site-api/internal/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Need to be imported for sqlite to make it work
	"github.com/stretchr/testify/suite"
)

// RepositoryTestSuite wrap logic around setup test for repository
type RepositoryTestSuite struct {
	suite.Suite
	Conn        *gorm.DB
	cleanupHook func()
}

// Setup prepare sqlite  database
func (s *RepositoryTestSuite) Setup() {
	log.Setup(&config.ApplicationConfiguration{
		LogToFile: false,
	})

	database.Connect(&config.DatabaseConfiguration{
		Driver:   "sqlite",
		Database: "/tmp/test.db",
		//Debug: true,
	})

	s.Conn = database.Instance()

}

// BeforeTest enable hook for cleaning database
func (s *RepositoryTestSuite) BeforeTest(suiteName, testName string) {
	log.Debugf("Before test %s from suite %s", suiteName, testName)
	s.cleanupHook = DeleteCreatedEntities(s.Conn)
}

// AfterTest trigger the hook
func (s *RepositoryTestSuite) AfterTest(suiteName, testName string) {
	log.Debugf("After test %s from suite %s", suiteName, testName)
	s.cleanupHook()
}
