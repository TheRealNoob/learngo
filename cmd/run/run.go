package run

import (
	"github.com/spf13/cobra"
	"github.com/therealnoob/learngo/cmd/run/scraper"
)

func Cmd() *cobra.Command {
	command := cobra.Command{
		Use:   "run",
		Short: "Run BotDetector component",
		Long:  "Run BotDetector component",
	}

	command.AddCommand(
		scraper.Cmd(),
	)

	return &command
}
