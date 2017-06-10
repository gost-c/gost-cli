package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestLoginCommand_implement(t *testing.T) {
	var _ cli.Command = &LoginCommand{}
}
