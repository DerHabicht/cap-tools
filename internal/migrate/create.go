package migrate

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/config"
	"github.com/derhabicht/cap-tools/internal/database"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the database in the config if it doesn't exist",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := create()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to create database")
		}

		os.Exit(0)
	},
}

func create() error {
	_, err := database.DBVersion()
	if err != nil {
		e := errors.Cause(err)
		if !strings.Contains(
			e.Error(),
			fmt.Sprintf(`database "%s" does not exist`, viper.GetString(config.DatabaseName)),
		) {
			return errors.WithStack(err)
		}
	} else {
		return nil
	}

	db, err := rootDB()
	if err != nil {
		return errors.WithStack(err)
	}
	defer db.Close()

	log.Info().Str("database", viper.GetString(config.DatabaseName)).Msg("creating database")
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", viper.GetString(config.DatabaseName)))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)
}
