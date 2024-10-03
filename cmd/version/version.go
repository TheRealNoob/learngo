package version

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	command := cobra.Command{
		Use:   "version",
		Short: "Print the version and exit",
		Long:  "Print the version and exit",
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf(
				"learngo Version: %s\nGo Version: %s\nGo OS/ARCH: %s %s\n",
				Version(),
				runtime.Version(),
				runtime.GOOS,
				runtime.GOARCH,
			)
		},
	}

	return &command
}

func Version() string {
	if env_version := os.Getenv("API_VERSION"); len(env_version) != 0 {
		return env_version
	} else {
		return "dev"
	}
}
