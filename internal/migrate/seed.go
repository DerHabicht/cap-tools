package migrate

import (
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/config"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := seed()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to seed database")
		}

		os.Exit(0)
	},
}

func seed() error {
	_, polluter, err := setup()
	if err != nil {
		return errors.WithStack(err)
	}

	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("seeding database")
	seedPath := viper.GetString(config.DatabaseMigrationSeed)

	f, err := os.Open(seedPath)
	if err != nil {
		log.Error().Err(err).Str("database", viper.GetString("database.name")).Msg("unable to read database seed file")
	}
	defer f.Close()

	err = polluter.Pollute(f)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
