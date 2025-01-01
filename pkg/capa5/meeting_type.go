package capa5

import (
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
