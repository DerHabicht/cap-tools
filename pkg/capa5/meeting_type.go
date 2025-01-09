package capa5

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type MT int

const (
	RegularMeeting MT = iota
	SpecialActivity
)

func ParseMT(s string) (MT, error) {
	switch strings.ToLower(s) {
	case "regular meeting":
		return RegularMeeting, nil
	case "special activity":
		return SpecialActivity, nil
	default:
		return -1, errors.Errorf("unable to parse meeting type: %q", s)
	}
}

func (mt MT) String() string {
	switch mt {
	case RegularMeeting:
		return "regular meeting"
	case SpecialActivity:
		return "special activity"
	default:
		panic(errors.Errorf("invalid meeting type: %d", mt))
	}
}

func (mt MT) MarshalJSON() ([]byte, error) {
	return []byte(mt.String()), nil
}

func (mt *MT) UnmarshalJSON(b []byte) error {
	s, err := ParseMT(string(b))
	if err != nil {
		return errors.WithStack(err)
	}

	*mt = s

	return nil
}

func (mt MT) Value() (driver.Value, error) {
	return mt.String(), nil
}

func (mt *MT) Scan(src interface{}) error {
	s, ok := src.(string)
	if !ok {
		return errors.Errorf("failed to scan '%v' into type %T", s, *mt)
	}

	p, err := ParseMT(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*mt = p

	return nil
}
