package pkg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMemberName(t *testing.T) {
	name := "John Jacob Jingleheimer Schmidt"

	first, middle, last := ParseMemberName(name)

	assert.Equal(t, "John", first)
	assert.Equal(t, "Jacob Jingleheimer", middle)
	assert.Equal(t, "Schmidt", last)

	name = "Robert Herschel Hawk"

	first, middle, last = ParseMemberName(name)

	assert.Equal(t, "Robert", first)
	assert.Equal(t, "Herschel", middle)
	assert.Equal(t, "Hawk", last)
}
