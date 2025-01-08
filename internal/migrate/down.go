package migrate

import (
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/config"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "run all down migrations against the database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := down()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to run down migrations")
		}

		os.Exit(0)
	},
}

func down() error {
	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("migrating database down")
	m, _, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Down()
	if err != nil {
		if !strings.Contains(err.Error(), "no migration") {
			return errors.WithStack(err)
		}
	}

	version, dirty, err := m.Version()
	if err != nil {
		if strings.Contains(err.Error(), "no migration") {
			log.Info().Str("database", viper.GetString(config.DatabaseName)).Uint("version", 0).Bool("dirty",
				false).Msg("down migration complete")
			return nil
		}
		return errors.WithStack(err)
	}

	log.Info().Str("database", viper.GetString(config.DatabaseName)).Uint("version", version).Bool("dirty",
		dirty).Msg("down migration complete")
	return nil
}

func init() {
	rootCmd.AddCommand(downCmd)
}
