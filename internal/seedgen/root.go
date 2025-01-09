package migrate

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "seedgen",
	Short: "Tool for generating seed data for dev/testing ",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&rootDBUser, "root-user", "u", "postgres", "root database user FOR DEV USE ONLY!")
	rootCmd.PersistentFlags().StringVarP(&rootDBPassword, "root-password", "p", "postgres", "root database password FOR DEV USE ONLY!")
	rootCmd.PersistentFlags().BoolVarP(&pollute, "seed", "s", false, "seed database")
}
