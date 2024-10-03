package run

import (
	"github.com/spf13/cobra"
	"github.com/therealnoob/learngo/cmd/run/scraper"
)

var (
	Command = &cobra.Command{
		Use:   "run",
		Short: "Run BotDetector component",
		Long:  "Run BotDetector component",
	}
)

func init() {
	Command.AddCommand(
		scraper.Command,
	)
}
