package command

import (
	"bufio"
	"flag"
	"github.com/mitchellh/cli"
	"io"
)

// Meta contain the meta-option that nearly all subcommand inherited.
type Meta struct {
	Ui cli.Ui
}

// NewFlagSet generates commom flag.FlagSet
func (m *Meta) NewFlagSet(name string, helpText string) *flag.FlagSet {
	flags := flag.NewFlagSet(name, flag.ContinueOnError)

	// Set usage function
	flags.Usage = func() { m.Ui.Error(helpText) }

	// Set error output to Meta.UI.Error
	errR, errW := io.Pipe()
	errScanner := bufio.NewScanner(errR)
	flags.SetOutput(errW)

	go func() {
		for errScanner.Scan() {
			m.Ui.Error(errScanner.Text())
		}
	}()

	return flags
}
