package main

import (
	"os"

	"github.com/gost-c/gost-cli/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}
