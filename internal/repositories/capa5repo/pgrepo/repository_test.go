package pgrepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/derhabicht/cap-tools/internal/database"
	"github.com/derhabicht/cap-tools/internal/repositories/capa5repo"
)

func TestPGRepoIsA5Repo(t *testing.T) {
	assert.NotPanics(t, func() {
		var _ capa5repo.A5Repository = NewPGRepository(nil)
	})
}

type PGRepoTestSuite struct {
	suite.Suite
	repo *PGRepository
}

func (s *PGRepoTestSuite) SetupSuite() {
	db, err := database.GetDB()
	assert.NoError(s.T(), err)
	s.repo = NewPGRepository(db)
}

func (s *PGRepoTestSuite) TearDownSuite() {
	err := s.repo.Close()
	assert.NoError(s.T(), err)
}

func (s *PGRepoTestSuite) SetupTest() {
	_, err := s.repo.db.Exec("DELETE FROM members;")
	assert.NoError(s.T(), err)

	_, err = s.repo.db.Exec("DELETE FROM units;")
	assert.NoError(s.T(), err)
}

func TestPGRepoTestSuite(t *testing.T) {
	suite.Run(t, new(PGRepoTestSuite))
}
