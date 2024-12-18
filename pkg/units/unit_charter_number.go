package units

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
)

var charterRegex = regexp.MustCompile(`(\w+)-(\w+)-(\d{3})`)

type UnitCharterNumber struct {
	region     Region
	wing       Wing
	unitNumber uint
}

func ParseCharterNumber(charterNumber string) (UnitCharterNumber, error) {
	res := charterRegex.FindStringSubmatch(charterNumber)
	if len(res) < 4 {
		return UnitCharterNumber{}, errors.Errorf("failed to parse charter number: %s", charterNumber)
	}
	region, err := ParseRegion(res[1])
	if err != nil {
		return UnitCharterNumber{}, errors.WithStack(err)
	}
	wing, err := ParseWing(res[2])
	if err != nil {
		return UnitCharterNumber{}, errors.WithStack(err)
	}
	unitNumber, err := strconv.Atoi(res[3])
	if err != nil {
		return UnitCharterNumber{}, errors.WithMessagef(err, "failed to parse units number from charter: %s", charterNumber)
	}

	return UnitCharterNumber{
		region:     region,
		wing:       wing,
		unitNumber: uint(unitNumber),
	}, nil
}

func (h UnitCharterNumber) Region() Region {
	return h.region
}

func (h UnitCharterNumber) Wing() Wing {
	return h.wing
}

func (h UnitCharterNumber) UnitNumber() uint {
	return h.unitNumber
}

func (h UnitCharterNumber) FullCharterNumber() string {
	return fmt.Sprintf("%s-%s-%03d", h.region, h.wing, h.unitNumber)
}

func (h UnitCharterNumber) ShortCharterNumber() string {
	return fmt.Sprintf("%s-%03d", h.wing, h.unitNumber)
}

func (h UnitCharterNumber) String() string {
	return h.FullCharterNumber()
}
