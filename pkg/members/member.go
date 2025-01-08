package members

import (
	"github.com/derhabicht/cap-tools/pkg/units"
)

type Member struct {
	capid          uint
	memberCategory MemberCategory
	lastName       string
	firstName      string
	grade          Grade
	unit           units.Unit
}

func NewMember(
	capid uint,
	memberCategory MemberCategory,
	lastName string,
	firstName string,
	grade Grade,
	unit units.Unit,
) Member {
	return Member{
		capid:          capid,
		memberCategory: memberCategory,
		lastName:       lastName,
		firstName:      firstName,
		grade:          grade,
		unit:           unit,
	}
}

func (m Member) CAPID() uint {
	return m.capid
}

func (m Member) MemberCategory() MemberCategory {
	return m.memberCategory
}

func (m Member) LastName() string {
	return m.lastName
}

func (m Member) FirstName() string {
	return m.firstName
}

func (m Member) Grade() Grade {
	return m.grade
}

func (m Member) Unit() units.Unit {
	return m.unit
}
