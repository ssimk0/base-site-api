package uploads

import (
	"base-site-api/models"
	"base-site-api/modules"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UploadTestSuite struct {
	modules.RepositoryTestSuite
}

func (s *UploadTestSuite) SetupTest() {
	s.Setup()
	s.Conn.Debug().AutoMigrate(
		&models.UploadType{},
		&models.UploadCategory{},
		&models.Upload{},
	)
}

func (s *UploadTestSuite) getTestUpload() *models.Upload {
	return &models.Upload{
		File:        "file",
		Thumbnail:   "thum",
		Description: "desc",
	}
}

func (s *UploadTestSuite) getTestUploadCategory() *models.UploadCategory {
	return &models.UploadCategory{
		Name: "Test",
		Slug: "test",
	}
}

func (s *UploadTestSuite) prepareTestData() ([]*models.Upload, []*models.UploadCategory, *models.UploadType) {
	uploadtype := &models.UploadType{
		Name: "media",
		Slug: "media",
	}
	if err := s.Conn.Create(uploadtype).Error; err != nil {
		panic(err)
	}

	uploads := []*models.Upload{
		{
			File:        "file",
			Thumbnail:   "thum",
			Description: "desc",
		},
		{
			File:        "file 2",
			Thumbnail:   "thum 2",
			Description: "desc 2",
		},
		{
			File:        "file 3 ",
			Thumbnail:   "thum 3 ",
			Description: "desc 3",
		},
	}

	uploadcategories := []*models.UploadCategory{
		{
			Name:    "Oznamy",
			Slug:    "oznamy",
			SubPath: "/oznamy",
		},
		{
			Name:    "Sluzby",
			Slug:    "sluzby",
			SubPath: "/sluzby",
		},
	}

	for _, a := range uploadcategories {
		a.TypeID = uploadtype.ID
		if err := s.Conn.Create(a).Error; err != nil {
			panic(err)
		}
	}

	for _, a := range uploads {
		a.CategoryID = uploadcategories[0].ID
		if err := s.Conn.Create(a).Error; err != nil {
			panic(err)
		}
	}

	return uploads, uploadcategories, uploadtype
}

func (s *UploadTestSuite) TestFindUploadsCategoryByTypes() {
	_, c, u := s.prepareTestData()
	r := NewRepository(s.Conn)

	category, err := r.FindCategoriesByType(u.Slug)

	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}

	assert.Len(s.T(), category, len(c))
	assert.Equal(s.T(), c[0].Name, category[0].Name)
}

func (s *UploadTestSuite) TestFindUploadsByCategory() {
	u, c, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	uploads, err := r.FindUploadsByCategory(c[0].Slug)
	if err != nil {
		s.T().Errorf("Error find page category by slug %s", err)
	}

	assert.Len(s.T(), uploads, len(u))
	assert.Equal(s.T(), u[0].File, uploads[0].File)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UploadTestSuite))
}
