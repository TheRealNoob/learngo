package scraper

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	httprest "github.com/therealnoob/learngo/pkg/http/rest"
)

func Cmd() *cobra.Command {
	command := cobra.Command{
		Use:   "scraper",
		Short: "API proxy for scraping data",
		Long:  "Scrapes Runescape Hiscores & RuneMetrics for player data.",
		Args:  cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			loadConfiguration(viper.GetString("config-file"))

			fmt.Printf("running scraper with config: %s\n", viper.GetString("config-file"))
			fmt.Printf("Host: %s\nPort: %d\n", viper.GetString("HttpServer.Host"), viper.GetUint16("HttpServer.Port"))

			handler := httprest.NewHandler()
			host := viper.GetString("httpServer.host") + ":" + viper.GetString("httpServer.port")

			fmt.Println("starting webserver...")
			if viper.GetBool("httpServer.TLS.enabled") {
				return http.ListenAndServeTLS(host, viper.GetString("httpServer.TLS.cert"), viper.GetString("httpServer.TLS.key"), handler)
			} else {
				return http.ListenAndServe(host, handler)
			}
		},
	}

	// define flags
	command.Flags().StringP("config-file", "c", "/etc/learngo/config.yaml", "path to config file")
	command.Flags().Bool("debug", false, "enable verbose output")

	// bind all pflags to viper
	viper.BindPFlags(command.Flags())

	// parse env vars
	viper.AutomaticEnv()

	return &command
}
