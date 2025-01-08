package units

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type UnitKind int

const (
	UnknownUnitKind UnitKind = iota
	NatlHQ
	RegionHQ
	WingHQ
	Group
	Squadron
	Flight
)

func ParseUnitKind(s string) UnitKind {
	switch strings.ToLower(s) {
	case "nhq":
		return NatlHQ
	case "region":
		return RegionHQ
	case "wing":
		return WingHQ
	case "group":
		return Group
	case "squadron":
		return Squadron
	case "flight":
		return Flight
	default:
		return UnknownUnitKind
	}
}

func (u UnitKind) String() string {
	switch u {
	case NatlHQ:
		return "nhq"
	case RegionHQ:
		return "region"
	case WingHQ:
		return "wing"
	case Group:
		return "group"
	case Squadron:
		return "squadron"
	case Flight:
		return "flight"
	default:
		return ""
	}
}

func (u UnitKind) MarshalJSON() ([]byte, error) {
	return []byte(u.String()), nil
}

func (u *UnitKind) UnmarshalJSON(b []byte) error {
	s := ParseUnitKind(string(b))

	*u = s

	return nil
}

func (u UnitKind) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *UnitKind) Scan(src any) error {
	s, ok := src.(string)
	if !ok {
		return errors.Errorf("failed to scan '%v' into type %T", s, *u)
	}

	p := ParseUnitKind(s)

	*u = p

	return nil
}
