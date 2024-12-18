package pkg

import (
	"strings"

	"github.com/pkg/errors"
)

type ContactType int

const (
	Email ContactType = iota
	CellPhone
	HomePhone
	WorkPhone
	CadetParentEmail
	CadetParentPhone
)

func ParseContactType(s string) (ContactType, error) {
	switch strings.ToLower(s) {
	case "email":
		return Email, nil
	case "cell phone":
		return CellPhone, nil
	case "home phone":
		return HomePhone, nil
	case "work phone":
		return WorkPhone, nil
	case "cadet parent email":
		return CadetParentEmail, nil
	case "cadet parent phone":
		return CadetParentPhone, nil
	default:
		return -1, errors.Errorf("invalid contact type: %s", s)
	}
}
