package config

import (
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	defaultIgnores := []string{"RMR-UT-000", "RMR-UT-999"}

	viper.SetDefault(UnitExclude, defaultIgnores)

	cfgDir, err := ConfigDir()
	if err != nil {
		log.Error().Err(err).Msg("could not find user config dir")
	} else {
		log.Debug().Str("usrCfgDir", cfgDir).Msg("user config dir found")
	}

	viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(cfgDir)
	err = viper.ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir, "config.yaml")
		log.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = viper.WriteConfigAs(path)
		if err != nil {
			log.Error().Err(err).Msg("failed to create default config")
		}
	}

}
