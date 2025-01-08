package migrate

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/config"
)

var dropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop the configured database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := drop()
		if err != nil {
			log.Fatal().Err(err).Msg("drop failed")
		}

		os.Exit(0)
	},
}

func drop() error {
	db, err := rootDB()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("dropping database")
	_, err = db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS %s;", viper.GetString(config.DatabaseName)))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(dropCmd)
}
