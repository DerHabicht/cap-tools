package units

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type UnitCategory int

const (
	UnknownUnitCategory UnitCategory = iota
	AdminUnit
	CompositeUnit
	CadetUnit
	SeniorUnit
)

func ParseUnitCategory(s string) UnitCategory {
	switch strings.ToLower(s) {
	case "admin":
		return AdminUnit
	case "composite":
		return CompositeUnit
	case "cadet":
		return CadetUnit
	case "senior":
		return SeniorUnit
	default:
		return UnknownUnitCategory
	}
}

func (u UnitCategory) String() string {
	switch u {
	case AdminUnit:
		return "admin"
	case CompositeUnit:
		return "composite"
	case CadetUnit:
		return "cadet"
	case SeniorUnit:
		return "senior"
	default:
		return ""
	}
}

func (u UnitCategory) MarshalJSON() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *UnitCategory) UnmarshalJSON(b []byte) error {
	s := ParseUnitCategory(string(b))

	*u = s

	return nil
}

func (u UnitCategory) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *UnitCategory) Scan(src any) error {
	s, ok := src.(string)
	if !ok {
		return errors.Errorf("failed to scan '%v' into type %T", s, *u)
	}

	p := ParseUnitCategory(s)

	*u = p

	return nil
}
