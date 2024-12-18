package config

import (
	"os"
	"path/filepath"

	"github.com/ag7if/cap/v2"
	set "github.com/deckarep/golang-set/v2"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	Exclude     = "exclude"
	UnitExclude = "unit_exclude"
)

var excludedEmails set.Set[string]
var excludedUnits set.Set[cap.UnitCharterNumber]

func ConfigDir() (string, error) {
	hd, err := os.UserConfigDir()
	if err != nil {
		return "", errors.WithMessage(err, "could not determine user config dir")
	}

	return filepath.Join(hd, "cap-tools"), nil
}

func ExcludedEmails() set.Set[string] {
	if excludedEmails == nil {
		emails := viper.GetStringSlice(Exclude)

		for _, email := range emails {
			excludedEmails.Add(email)
		}
	}

	return excludedEmails
}

func ExcludedUnits() set.Set[cap.UnitCharterNumber] {
	if excludedUnits == nil {
		units := viper.GetStringSlice(UnitExclude)

		for _, unit := range units {
			cn, err := cap.ParseCharterNumber(unit)
			if err != nil {
				log.Warn().Err(err).Str("charter_number", unit).Msg("could not parse charter number")
				continue
			}
			excludedUnits.Add(cn)
		}
	}

	return excludedUnits
}
