package migrate

import (
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Run all down migrations, then run all up migrations",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := reset()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to run reset")
		}

		os.Exit(1)
	},
}

func reset() error {
	err := down()
	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		return errors.WithStack(err)
	}
	err = up()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
