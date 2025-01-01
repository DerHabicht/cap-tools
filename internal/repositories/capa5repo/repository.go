package capa5repo

import (
	"github.com/google/uuid"

	"github.com/derhabicht/cap-tools/pkg/capa5"
	"github.com/derhabicht/cap-tools/pkg/members"
	"github.com/derhabicht/cap-tools/pkg/units"
)

type A5Repository interface {
	CreateUnit(unit units.Unit) error
	ListUnits() []units.Unit
	FetchUnit(charter units.UnitCharterNumber) (*units.Unit, error)
	UpdateUnit(unit units.Unit) error
	DeleteUnit(unit units.Unit) error

	CreateMeeting(unit capa5.Meeting) error
	ListMeetings(charter units.UnitCharterNumber) []capa5.Meeting
	FetchMeeting(id uuid.UUID) (*units.Unit, error)
	UpdateMeeting(meeting capa5.Meeting) error
	DeleteMeeting(meeting capa5.Meeting) error

	CreateMember(member members.Member) error
	ListMembers(charter units.UnitCharterNumber) []members.Member
	FetchMember(capid uint) (members.Member, error)
	UpdateMember(member members.Member) error
	DeleteMember(member members.Member) error
}
