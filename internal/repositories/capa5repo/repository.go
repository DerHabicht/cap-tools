package capa5repo

import (
	"github.com/google/uuid"

	"github.com/derhabicht/cap-tools/pkg/capa5"
	"github.com/derhabicht/cap-tools/pkg/members"
	"github.com/derhabicht/cap-tools/pkg/units"
)

type A5Repository interface {
	CreateOrUpdateUnit(unit units.Unit) error
	ListUnits(searchName string, offset, limit int) ([]units.Unit, error)
	FetchUnit(charterNumber units.UnitCharterNumber) (units.Unit, error)
	DeleteUnit(unit units.Unit) error

	CreateOrUpdateMeeting(meeting capa5.Meeting) error
	ListMeetings(charterNumber units.UnitCharterNumber, offset, limit int) ([]capa5.Meeting, error)
	FetchMeeting(meetingID uuid.UUID) (capa5.Meeting, error)
	DeleteMeeting(meeting capa5.Meeting) error

	CreateOrUpdateMember(member members.Member) error
	ListMembers(charterNumber units.UnitCharterNumber, offset, limit int) ([]members.Member, error)
	FetchMember(capid uint) (members.Member, error)
	DeleteMember(member members.Member) error

	CreateOrUpdateUser(user capa5.User) error
	ListUsers(charterNumber units.UnitCharterNumber, offset, limit int) ([]capa5.User, error)
	FetchUser(id uuid.UUID) (capa5.User, error)
	DeleteUser(user capa5.User) error
}
