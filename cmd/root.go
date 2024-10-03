package root

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/therealnoob/learngo/cmd/run"
	"github.com/therealnoob/learngo/cmd/version"
)

func Cmd() *cobra.Command {
	command := cobra.Command{
		Use:   "learngo",
		Short: "Run BotDetector scraper",
		Long:  "Scrapes OldSchool Runescape HiScores & RuneMetrics",
	}

	command.AddCommand(
		run.Cmd(),
		version.Cmd(),
	)

	viper.SetEnvPrefix("LEARNGO")
	// replaces in pairs - both . and - become _
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	return &command
}
