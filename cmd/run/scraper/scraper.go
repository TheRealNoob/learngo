package scraper

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	Command = &cobra.Command{
		Use:   "scraper",
		Short: "API proxy for scraping data",
		Long:  "Scrapes Runescape Hiscores & RuneMetrics for player data.",
		Args:  cobra.MinimumNArgs(0),
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		Run: func(cmd *cobra.Command, args []string) {
			loadConfiguration(viper.GetString("config-file"))

			fmt.Printf("running scraper with config: %s\n", viper.GetString("config-file"))
			fmt.Printf("Host: %s\nPort: %d\n", viper.GetString("HttpServer.Host"), viper.GetUint16("HttpServer.Port"))
		},
	}
)

func init() {
	// define flags
	Command.Flags().StringP("config-file", "c", "/etc/learngo/config.yaml", "path to config file")
	Command.Flags().Bool("debug", false, "enable verbose output")

	// bind all pflags to viper
	viper.BindPFlags(Command.Flags())

	// parse env vars
	viper.AutomaticEnv()
}
