package command

import (
	"strings"
)

type LoginCommand struct {
	Meta
}

func (c *LoginCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *LoginCommand) Synopsis() string {
	return "Login your account"
}

func (c *LoginCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
