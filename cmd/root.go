package root

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/therealnoob/learngo/cmd/run"
	"github.com/therealnoob/learngo/cmd/version"
)

var (
	Command = &cobra.Command{
		Use:   "learngo",
		Short: "Run BotDetector scraper",
		Long:  "Scrapes OldSchool Runescape HiScores & RuneMetrics",
	}
)

func init() {
	Command.AddCommand(
		run.Command,
		version.Command,
	)

	viper.SetEnvPrefix("LEARNGO")
	// replaces in pairs - both . and - become _
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}
