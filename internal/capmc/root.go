package capmc

import (
	"os"

	"github.com/ag7if/go-files"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/derhabicht/cap-tools/internal/contact"
)

var rootCmd = &cobra.Command{
	Use:   "capmc <mbr_contact_report>",
	Short: "Ingest the given member contact report CSV and print just the emails to STDOUT.",
	Long: `Ingest the given member contact report CSV and print just the emails to STDOUT.

Unit with a charter number of 000 (reserve units) and 999 (
legislative units) are automatically excluded when scraping the report.
`,
	Args: cobra.ExactArgs(1),
	Run:  runRoot,
}

func runRoot(cmd *cobra.Command, args []string) {
	reportFile, err := files.NewFile(args[0], &log.Logger)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open report file")
	}

	err = contact.ScrapeEmails(reportFile)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open report file")
	}
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
}
