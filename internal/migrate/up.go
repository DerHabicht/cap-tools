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

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Run all up migrations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := up()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to up migrations")
		}

		os.Exit(0)
	},
}

func up() error {
	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("migrating database up")
	m, _, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	err = m.Up()
	if err != nil {
		if !strings.Contains(err.Error(), "no change") {
			return errors.WithStack(err)
		}
	}

	version, dirty, err := m.Version()
	if err != nil {
		return errors.WithStack(err)
	}
	log.Info().Str("database", viper.GetString(config.DatabaseName)).Uint("version", version).Bool("dirty",
		dirty).Msg("up migration complete")

	if pollute {
		err = seed()
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(upCmd)
}
