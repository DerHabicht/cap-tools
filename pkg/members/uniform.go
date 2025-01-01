package members

import (
	"database/sql/driver"
	"strings"

	"github.com/pkg/errors"
)

type Uniform int

const (
	MessDress Uniform = iota
	SemiFormal
	ServiceDress
	ClassB
	ABU
	FDU
	PTU
	CorporateSemiFormal
	CorporateServiceDress
	Aviator
	CWU
	CFU
	CFDU
	CivilianAttire
)

func ParseUniform(s string) (Uniform, error) {
	switch strings.ToLower(s) {
	case "usaf mess dress":
		fallthrough
	case "messdress":
		return MessDress, nil
	case "usaf semi-formal":
		fallthrough
	case "semiformal":
		return SemiFormal, nil
	case "usaf service dress":
		fallthrough
	case "servicedress":
		return ServiceDress, nil
	case "usaf class b":
		fallthrough
	case "classb":
		return ClassB, nil
	case "abu":
		return ABU, nil
	case "fdu":
		return FDU, nil
	case "ptu":
		return PTU, nil
	case "corporate semi-formal":
		fallthrough
	case "corporatesemiformal":
		return CorporateSemiFormal, nil
	case "corporate service dress":
		fallthrough
	case "corporateservicedress":
		return CorporateServiceDress, nil
	case "aviator combination":
		fallthrough
	case "aviator":
		return Aviator, nil
	case "cwu":
		return CWU, nil
	case "cfu":
		return CFU, nil
	case "cfdu":
		return CFDU, nil
	case "civilian attire":
		fallthrough
	case "civilianattire":
		return CivilianAttire, nil
	default:
		return -1, errors.Errorf("unrecognized uniform: %s", s)
	}
}

func (u Uniform) String() string {
	switch u {
	case MessDress:
		return "MessDress"
	case SemiFormal:
		return "SemiFormal"
	case ServiceDress:
		return "ServiceDress"
	case ClassB:
		return "ClassB"
	case ABU:
		return "ABU"
	case FDU:
		return "FDU"
	case PTU:
		return "PTU"
	case CorporateSemiFormal:
		return "CorporateSemiFormal"
	case CorporateServiceDress:
		return "CorporateServiceDress"
	case Aviator:
		return "Aviator"
	case CWU:
		return "CWU"
	case CFU:
		return "CFU"
	case CFDU:
		return "CFDU"
	case CivilianAttire:
		return "CivilianAttire"
	default:
		panic(errors.Errorf("invalid uniform value: %d", u))
	}
}

func (u Uniform) MarshalJSON() ([]byte, error) {
	switch u {
	case MessDress:
		return []byte("USAF Mess Dress"), nil
	case SemiFormal:
		return []byte("USAF Semi-Formal"), nil
	case ServiceDress:
		return []byte("USAF Service Dress"), nil
	case ClassB:
		return []byte("USAF Class B"), nil
	case ABU:
		return []byte("ABU"), nil
	case FDU:
		return []byte("FDU"), nil
	case PTU:
		return []byte("PTU"), nil
	case CorporateSemiFormal:
		return []byte("Corporate Semi-Formal"), nil
	case CorporateServiceDress:
		return []byte("Corporate Service Dress"), nil
	case Aviator:
		return []byte("Corporate Aviator"), nil
	case CWU:
		return []byte("CWU"), nil
	case CFU:
		return []byte("CFU"), nil
	case CFDU:
		return []byte("CFDU"), nil
	case CivilianAttire:
		return []byte("Civilian Attire"), nil
	default:
		return nil, errors.Errorf("invalid uniform value: %d", u)
	}
}

func (u *Uniform) UnmarshalJSON(b []byte) error {
	s, err := ParseUniform(string(b))
	if err != nil {
		return errors.WithStack(err)
	}

	*u = s

	return nil
}

func (u Uniform) Value() (driver.Value, error) {
	return u.String(), nil
}

func (u *Uniform) Scan(src any) error {
	s, ok := src.(string)
	if !ok {
		return errors.Errorf("failed to scan value as type Uniform: %v", src)
	}

	p, err := ParseUniform(s)
	if err != nil {
		return errors.WithStack(err)
	}

	*u = p

	return nil
}
