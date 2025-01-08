package pgrepo

import (
	"github.com/stretchr/testify/assert"

	"github.com/derhabicht/cap-tools/pkg/members"
	"github.com/derhabicht/cap-tools/pkg/units"
)

func (s *PGRepoTestSuite) TestMembers() {
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

	res, err := s.repo.FetchMember(capid)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), member.CAPID(), res.CAPID())
	assert.Equal(s.T(), member.LastName(), res.LastName())
	assert.Equal(s.T(), member.FirstName(), res.FirstName())
	assert.Equal(s.T(), member.Grade(), res.Grade())
	assert.Equal(s.T(), member.Unit().CharterNumber(), res.Unit().CharterNumber())

	member = members.NewMember(
		member.CAPID(),
		member.MemberCategory(),
		"Malfunction",
		member.FirstName(),
		members.Maj,
		member.Unit(),
	)

	err = s.repo.CreateOrUpdateMember(member)
	assert.NoError(s.T(), err)

	res, err = s.repo.FetchMember(capid)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), member.CAPID(), res.CAPID())
	assert.Equal(s.T(), member.LastName(), res.LastName())
	assert.Equal(s.T(), member.FirstName(), res.FirstName())
	assert.Equal(s.T(), member.Grade(), res.Grade())

	list, err := s.repo.ListMembers(cn, 0, 10)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), len(list), 1)

	err = s.repo.DeleteMember(member)
	assert.NoError(s.T(), err)

	res, err = s.repo.FetchMember(capid)
	assert.ErrorAs(s.T(), err, &ErrNotFound{})
}
