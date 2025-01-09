package capa5

import (
	"github.com/google/uuid"

	"github.com/derhabicht/cap-tools/pkg/units"
)

type User struct {
	id            uuid.UUID
	username      string
	roles         map[units.UnitCharterNumber]Role
	linkedMembers []MemberLink
}

func NewUser(
	id uuid.UUID,
	username string,
	roles map[units.UnitCharterNumber]Role,
	linkedMembers []MemberLink,
) User {
	return User{
		id:            id,
		username:      username,
		roles:         roles,
		linkedMembers: linkedMembers,
	}
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) Username() string {
	return u.username
}

func (u User) Roles() map[units.UnitCharterNumber]Role {
	return u.roles
}

func (u User) LinkedMembers() []MemberLink {
	return u.linkedMembers
}
