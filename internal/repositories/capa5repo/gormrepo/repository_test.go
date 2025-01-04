package gormrepo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/derhabicht/cap-tools/internal/repositories/capa5repo"
)

func TestGormRepoIsA5Repo(t *testing.T) {
	repo := GormRepository{}
	var i any = repo
	_, ok := i.(capa5repo.A5Repository)

	assert.True(t, ok)
}
