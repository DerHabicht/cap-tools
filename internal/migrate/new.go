package migrate

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/config"
)

var newCmd = &cobra.Command{
	Use:   "new [name]",
	Short: "Create a new migration",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := newMigration(args[0])
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to create a new migration")
		}

		os.Exit(0)
	},
}

func newMigration(name string) error {
	path := viper.GetString(config.DatabaseMigrationSource)[7:]

	dir := files.NewDirectory(path)
	err := dir.Populate()
	if err != nil {
		return errors.WithStack(err)
	}

	migrations := dir.FilterByExt("sql")

	// Since Filter returns a sorted list, the last file in the list should be the highest number we have
	verStr := migrations[len(migrations)-1].Name()
	ver, err := strconv.Atoi(strings.SplitN(verStr, "_", 2)[0])
	if err != nil {
		return errors.WithStack(err)
	}

	ver++

	_, err = dir.CreateFile(fmt.Sprintf("%06d_%s.up.sql", ver, name))
	if err != nil {
		return errors.WithStack(err)
	}

	_, err = dir.CreateFile(fmt.Sprintf("%06d_%s.down.sql", ver, name))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(newCmd)
}
