package members

import (
	"strings"

	"github.com/pkg/errors"
)

type MemberType int

const (
	Senior MemberType = iota
	Cadet
	CadetSponsor
	AEM
	StateLegislative
	Legislative
	Patron
)

func ParseMemberType(s string) (MemberType, error) {
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

func (m MemberType) String() string {
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

	panic(errors.Errorf("invalid MemberType value: %d", m))
}
