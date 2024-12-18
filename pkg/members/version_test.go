package members

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/mod/semver"
)

func TestVersion(t *testing.T) {
	v := Version()
	assert.True(t, semver.IsValid(v))
}
