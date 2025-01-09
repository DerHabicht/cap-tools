package pgrepo

import (
	"github.com/stretchr/testify/assert"

	"github.com/derhabicht/cap-tools/pkg/members"
	"github.com/derhabicht/cap-tools/pkg/units"
)

func (s *PGRepoTestSuite) TestUsers() {
	cn, err := units.ParseCharterNumber("RMR-UT-080")
	assert.NoError(s.T(), err)

	unit := units.NewUnit(
		cn,
		units.Squadron,
		units.CadetUnit,
		"Blackhawk Cadet Squadron",
		"12953 S Minuteman Dr",
		"Draper, UT 84020",
	)

	err = s.repo.CreateOrUpdateUnit(unit)
	assert.NoError(s.T(), err)

	var capid uint = 999999

	member := members.NewMember(
		capid,
		members.Senior,
		"Obvious",
		"Raleigh",
		members.Capt,
		unit,
	)

	err = s.repo.CreateOrUpdateMember(member)
	assert.NoError(s.T(), err)

	user :=
}
