package config

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func ConfigDir() (string, error) {
	hd, err := os.UserConfigDir()
	if err != nil {
		return "", errors.WithMessage(err, "could not determine user config dir")
	}

	return filepath.Join(hd, "cap-tools"), nil
}
