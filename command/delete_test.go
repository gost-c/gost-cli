package command

import (
	"github.com/mitchellh/cli"
	"testing"
)

func TestDeleteCommand_implement(t *testing.T) {
	var _ cli.Command = &DeleteCommand{}
}
