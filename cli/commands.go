package cli

import (
	"github.com/gost-c/gost-cli/command"
	"github.com/mitchellh/cli"
)

// Commands is a entry function to run sub commands
func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"register": func() (cli.Command, error) {
			return &command.RegisterCommand{
				Meta: *meta,
			}, nil
		},
		"login": func() (cli.Command, error) {
			return &command.LoginCommand{
				Meta: *meta,
			}, nil
		},
		"push": func() (cli.Command, error) {
			return &command.PushCommand{
				Meta: *meta,
			}, nil
		},
		"delete": func() (cli.Command, error) {
			return &command.DeleteCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta: *meta,
			}, nil
		},
	}
}
