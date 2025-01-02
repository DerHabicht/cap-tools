package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

const (
	DatabaseHost            = "database.host"
	DatabasePort            = "database.port"
	DatabaseName            = "database.name"
	DatabaseUser            = "database.user"
	DatabasePassword        = "database.password"
	DatabaseSSL             = "database.ssl"
	DatabaseMigrationSource = "database.migration.source"
	DatabaseMigrationSeed   = "database.migration.seed"
)

const (
	appName       = "cap-tools"
	baseConfigKey = "CAPTOOLS"
)

func ConfigDir() (string, error) {
	cfgDir := os.Getenv(fmt.Sprintf("%s_CONFIG", baseConfigKey))
	if cfgDir != "" {
		return cfgDir, nil
	}

	hd, err := os.UserConfigDir()
	if err != nil {
		return "", errors.WithMessage(err, "could not determine user config dir")
	}

	return filepath.Join(hd, appName), nil
}

func CacheDir() (string, error) {
	cacheDir := os.Getenv(fmt.Sprintf("%s_CACHE", baseConfigKey))
	if cacheDir != "" {
		return cacheDir, nil
	}

	cd, err := os.UserCacheDir()
	if err != nil {
		return "", errors.WithMessage(err, "could not determine user cache dir")
	}

	return filepath.Join(cd, appName), nil
}
