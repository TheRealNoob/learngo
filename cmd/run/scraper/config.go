package scraper

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	httpServer httpServer
	httpClient httpClient
	logging    logging
}

type httpServer struct {
	host string
	port uint16
	TLS  struct {
		enabled    bool
		certFile   string
		keyFile    string
		minVersion float64
		maxVersion float64
	}
}

type httpClient struct {
	TLS struct {
		ignoreInvalidCert bool
	}
}

type logging struct {
	level  string
	format string
}

// func loadConfiguration(configFile string) *Configuration {
func loadConfiguration(configFile string) {
	var config Configuration

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
