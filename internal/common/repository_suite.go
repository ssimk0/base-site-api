package common

import (
	"base-site-api/internal/log"
	"base-site-api/internal/test_helper"

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

	var err error
	s.Conn, err = gorm.Open("sqlite3", "/tmp/gorm.db")
	if err != nil {
		log.Fatal(err)
	}

	// s.Conn.LogMode(true)
}

// BeforeTest enable hook for cleaning database
func (s *RepositoryTestSuite) BeforeTest(suiteName, testName string) {
	log.Debugf("Before test %s from suite %s", suiteName, testName)
	s.cleanupHook = test_helper.DeleteCreatedEntities(s.Conn)
}

// AfterTest trigger the hook
func (s *RepositoryTestSuite) AfterTest(suiteName, testName string) {
	log.Debugf("After test %s from suite %s", suiteName, testName)
	s.cleanupHook()
}
