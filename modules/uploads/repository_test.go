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

	uploads, count, err := r.FindUploadsByCategory(c[0].Slug, 0, 10)
	if err != nil {
		s.T().Errorf("Error find uploads category by slug %s", err)
	}

	assert.Len(s.T(), uploads, len(u))
	assert.Equal(s.T(), 3, count)
	assert.Equal(s.T(), u[0].File, uploads[0].File)
}

func (s *UploadTestSuite) TestUpdateUploadCategory() {
	_, c, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	err := r.UpdateCategory("New name", "updated", "", c[0].ID)
	if err != nil {
		s.T().Errorf("Error update upload %s", err)
	}
	category := models.UploadCategory{}
	s.Conn.Find(&category, c[0].ID)

	assert.Equal(s.T(), category.SubPath, "updated")
	assert.Equal(s.T(), category.Name, "New name")
}

func (s *UploadTestSuite) TestUpdateUpload() {
	u, _, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	err := r.Update("updated desc", u[0].ID)
	if err != nil {
		s.T().Errorf("Error update upload category %s", err)
	}
	upload := models.Upload{}
	s.Conn.Find(&upload, u[0].ID)

	assert.Equal(s.T(), upload.Description, "updated desc")
}

func (s *UploadTestSuite) TestStoreUploadCategory() {
	_, _, t := s.prepareTestData()
	r := NewRepository(s.Conn)

	c := models.UploadCategory{
		Name:   "file cat",
		Slug:   "file-cat",
		TypeID: t.ID,
	}
	id, err := r.StoreCategory(&c)
	if err != nil {
		s.T().Errorf("Error create upload category %s", err)
	}
	upload := models.UploadCategory{}
	s.Conn.Find(&upload, id)

	assert.Equal(s.T(), c.Name, upload.Name)
	assert.Equal(s.T(), c.TypeID, upload.TypeID)
	assert.NotEqual(s.T(), 0, id)
}

func (s *UploadTestSuite) TestStoreUpload() {
	_, c, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	u := models.Upload{
		File:        "file",
		Description: "desc",
		CategoryID:  c[1].ID,
	}
	id, err := r.Store(&u)
	if err != nil {
		s.T().Errorf("Error create upload %s", err)
	}
	upload := models.Upload{}
	s.Conn.Find(&upload, id)

	assert.Equal(s.T(), u.Description, upload.Description)
	assert.Equal(s.T(), u.File, upload.File)
	assert.NotEqual(s.T(), 0, id)
}

func (s *UploadTestSuite) TestDelete() {
	u, _, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	err := r.Delete(u[0].ID)

	if err != nil {
		s.T().Errorf("Error delete upload %s", err)
	}

	a := &models.Upload{}
	s.Conn.First(a, u[0].ID)
	// Not found
	assert.Equal(s.T(), uint(0), a.ID)
}

func (s *UploadTestSuite) TestDeleteCategory() {
	_, c, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	err := r.DeleteCategory(c[0].ID)

	if err != nil {
		s.T().Errorf("Error delete upload category %s", err)
	}

	a := &models.Upload{}
	s.Conn.First(a, c[0].ID)
	// Not found
	assert.Equal(s.T(), uint(0), a.ID)
}

func (s *UploadTestSuite) TestFindUpload() {
	u, _, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	a, err := r.Find(u[0].ID)

	if err != nil {
		s.T().Errorf("Error find upload %s", err)
	}

	// found
	assert.Equal(s.T(), u[0].ID, a.ID)
}

func (s *UploadTestSuite) TestFindTypeBySlug() {
	_, _, t := s.prepareTestData()
	r := NewRepository(s.Conn)

	a, err := r.FindTypeBySlug(t.Slug)

	if err != nil {
		s.T().Errorf("Error find type upload %s", err)
	}

	// found
	assert.NotEqual(s.T(), 0, t.ID)
	assert.Equal(s.T(), t.ID, a.ID)
}

func (s *UploadTestSuite) TestFindUploadCategoryBySlug() {
	_, c, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	a, err := r.FindCategoryBySlug(c[0].Slug)

	if err != nil {
		s.T().Errorf("Error find upload category by slug %s", err)
	}

	// found
	assert.Equal(s.T(), c[0].ID, a.ID)
}

func (s *UploadTestSuite) TestFindUploadCategory() {
	_, c, _ := s.prepareTestData()
	r := NewRepository(s.Conn)

	a, err := r.FindCategory(c[0].ID)

	if err != nil {
		s.T().Errorf("Error find upload category %s", err)
	}

	// found
	assert.Equal(s.T(), c[0].ID, a.ID)
}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UploadTestSuite))
}
