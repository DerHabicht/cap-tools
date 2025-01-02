package capa5

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "capa5srv",
	Short: "Run the CAP/A5 server",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	Run:   runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
}

func Execute(version string) {
	viper.Set("version", version)
	rootCmd.Version = version
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("loglevel", "", "info", "Log level")
	err := viper.BindPFlag("loglevel", rootCmd.PersistentFlags().Lookup("loglevel"))
	if err != nil {
		log.Error().Err(err).Msg("failed to bind log level")
	}

	rootCmd.Flags().UintP("port", "p", 8080, "Port to listen on")
	err = viper.BindPFlag("port", rootCmd.Flags().Lookup("port"))
}
