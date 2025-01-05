package capa5

import (
	"github.com/google/uuid"
)

type MemberLink struct {
	userID      uuid.UUID
	memberCAPID uint
	self        bool
}

func NewMemberLink(
	userID uuid.UUID,
	memberCAPID uint,
	self bool,
) MemberLink {
	return MemberLink{
		userID:      userID,
		memberCAPID: memberCAPID,
		self:        self,
	}
}

func (m MemberLink) UserID() uuid.UUID {
	return m.userID
}

func (m MemberLink) MemberCAPID() uint {
	return m.memberCAPID
}

func (m MemberLink) Self() bool {
	return m.self
}
