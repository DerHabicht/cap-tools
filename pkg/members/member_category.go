package members

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type MemberCategory int

const (
	Senior MemberCategory = iota
	Cadet
	CadetSponsor
	AEM
	StateLegislative
	Legislative
	Patron
)

func ParseMemberCategory(s string) (MemberCategory, error) {
	switch strings.ToLower(s) {
	case "senior":
		return Senior, nil
	case "cadet":
		return Cadet, nil
	case "cadet sponsor":
		return CadetSponsor, nil
	case "aem":
		return AEM, nil
	case "state leg":
		return StateLegislative, nil
	case "legislative":
		return Legislative, nil
	case "patron":
		return Patron, nil
	default:
		return -1, errors.Errorf("unrecognized members type: %s", s)
	}
}

func (m MemberCategory) String() string {
	switch m {
	case Senior:
		return "SENIOR"
	case Cadet:
		return "CADET"
	case CadetSponsor:
		return "CADET SPONSOR"
	case AEM:
		return "AEM"
	case StateLegislative:
		return "STATE LEG"
	case Legislative:
		return "LEGISLATIVE"
	case Patron:
		return "PATRON"
	}

	panic(errors.Errorf("invalid MemberCategory value: %d", m))
}

func (m MemberCategory) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

func (m *MemberCategory) UnmarshalJSON(b []byte) error {
	s, err := ParseMemberCategory(string(b))
	if err != nil {
		return errors.WithStack(err)
	}

	*m = s

	return nil
}

func (m MemberCategory) Value() (driver.Value, error) {
	return m.String(), nil
}

func (m *MemberCategory) Scan(src interface{}) error {
	s, ok := src.(string)
	if !ok {
		return errors.Errorf("failed to scan '%v' into type %T", s, *m)
	}

	p, err := ParseMemberCategory(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*m = p

	return nil
}
