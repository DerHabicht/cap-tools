package pgrepo

import (
	"github.com/stretchr/testify/assert"

	"github.com/derhabicht/cap-tools/pkg/members"
)

func (s *PGRepoTestSuite) TestMembers() {
	var capid uint = 999999

	member := members.NewMember(
		capid,
		"Obvious",
		"Raleigh",
		members.Capt,
	)

	err := s.repo.CreateOrUpdateMember(member)
	assert.NoError(s.T(), err)

	res, err := s.repo.FetchMember(capid)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), member.CAPID(), res.CAPID())
	assert.Equal(s.T(), member.LastName(), res.LastName())
	assert.Equal(s.T(), member.FirstName(), res.FirstName())
	assert.Equal(s.T(), member.Grade(), res.Grade())

	member = members.NewMember(
		member.CAPID(),
		"Malfunction",
		member.FirstName(),
		members.Maj,
	)

	err = s.repo.CreateOrUpdateMember(member)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), member.CAPID(), res.CAPID())
	assert.Equal(s.T(), member.LastName(), res.LastName())
	assert.Equal(s.T(), member.FirstName(), res.FirstName())
	assert.Equal(s.T(), member.Grade(), res.Grade())

}
