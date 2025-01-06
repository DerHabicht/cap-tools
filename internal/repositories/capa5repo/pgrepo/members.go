package pgrepo

import (
	"github.com/derhabicht/cap-tools/pkg/members"
	"github.com/derhabicht/cap-tools/pkg/units"
)

func (repo *PGRepository) CreateOrUpdateMember(member members.Member) error {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) ListMembers(charterNumber units.UnitCharterNumber, offset, limit int) ([]members.Member, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) FetchMember(capid uint) (members.Member, error) {
	//TODO implement me
	panic("implement me")
}

func (repo *PGRepository) DeleteMember(member members.Member) error {
	//TODO implement me
	panic("implement me")
}
