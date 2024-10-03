package main

import (
	"fmt"
	"os"

	root "github.com/therealnoob/learngo/cmd"
)

func main() {
	if err := root.Cmd().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
