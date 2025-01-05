package pgrepo

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/derhabicht/cap-tools/internal/database"
	"github.com/derhabicht/cap-tools/internal/repositories/capa5repo"
)

func TestPGRepoIsA5Repo(t *testing.T) {
	repo := PGRepository{}
	var i any = repo
	_, ok := i.(capa5repo.A5Repository)

	assert.True(t, ok)
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

func TestPGRepoTestSuite(t *testing.T) {
	suite.Run(t, new(PGRepoTestSuite))
}
