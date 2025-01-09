package capa5

import (
	"database/sql/driver"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Role int

const (
	Root Role = iota
	Command
	Admin
	Member
	Parent
)

func ParseRole(s string) (Role, error) {
	switch strings.ToLower(s) {
	case "root":
		return Root, nil
	case "command":
		return Command, nil
	case "admin":
		return Admin, nil
	case "member":
		return Member, nil
	case "parent":
		return Parent, nil
	default:
		return -1, errors.Errorf("invalid role %q", s)
	}
}

func (r Role) String() string {
	switch r {
	case Root:
		return "ROOT"
	case Command:
		return "COMMAND"
	case Admin:
		return "ADMIN"
	case Member:
		return "MEMBER"
	case Parent:
		return "PARENT"
	default:
		panic(fmt.Sprintf("invalid role %d", r))
	}
}

func (r Role) MarshalJSON() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *Role) UnmarshalJSON(b []byte) error {
	s, err := ParseRole(string(b))
	if err != nil {
		return errors.WithStack(err)
	}

	*r = s

	return nil
}

func (r Role) Value() (driver.Value, error) {
	return r.String(), nil
}

func (r *Role) Scan(src interface{}) error {
	s, ok := src.(string)
	if !ok {
		return errors.Errorf("failed to scan '%v' into type %T", s, *r)
	}

	p, err := ParseRole(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*r = p

	return nil
}
