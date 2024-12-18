package cap

import (
	"strings"

	"github.com/pkg/errors"
)

type Region uint

const (
	NortheastRegion Region = 91 + iota
	MidAtlanticRegion
	GreatLakesRegion
	SoutheastRegion
	NorthCentralRegion
	SouthwestRegion
	RockyMountainRegion
	PacificRegion
	NationalHeadquarters
)

func ParseRegion(s string) (Region, error) {
	switch strings.ToUpper(s) {
	case "NER":
		return NortheastRegion, nil
	case "MAR":
		return MidAtlanticRegion, nil
	case "GLR":
		return GreatLakesRegion, nil
	case "SER":
		return SoutheastRegion, nil
	case "NCR":
		return NorthCentralRegion, nil
	case "SWR":
		return SouthwestRegion, nil
	case "RMR":
		return RockyMountainRegion, nil
	case "PCR":
		return PacificRegion, nil
	case "NHQ":
		return NationalHeadquarters, nil
	default:
		return 0, errors.Errorf("invalid region: %s", s)
	}
}

func (r Region) String() string {
	switch r {
	case NortheastRegion:
		return "NER"
	case MidAtlanticRegion:
		return "MAR"
	case GreatLakesRegion:
		return "GLR"
	case SoutheastRegion:
		return "SER"
	case NorthCentralRegion:
		return "NCR"
	case SouthwestRegion:
		return "SWR"
	case RockyMountainRegion:
		return "RMR"
	case PacificRegion:
		return "PCR"
	case NationalHeadquarters:
		return "NHQ"
	default:
		panic(errors.Errorf("invalid region code: %d", r))
	}
}
