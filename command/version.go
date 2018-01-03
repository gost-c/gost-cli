package command

import (
	"fmt"
)

// Name is cli's name
const Name = "gost"

// Version is cli's current version
const Version = "v1.2.1"

// GitCommit describes latest commit hash.
// This value is extracted by git command when building.
// To set this from outside, use go build -ldflags "-X main.GitCommit \"$(COMMIT)\""
var GitCommit string

// VersionCommand is struct of login meta
type VersionCommand struct {
	Meta
}

// Run is entry function of version command
func (c *VersionCommand) Run(args []string) int {
	version := fmt.Sprintf("\n%s version %s", Name, Version)
	if len(GitCommit) != 0 {
		version += fmt.Sprintf(" (%s)", GitCommit)
	}
	fmt.Println(version)
	return 0
}

// Synopsis is description of version command
func (c *VersionCommand) Synopsis() string {
	return fmt.Sprintf("Print %s version and quit", Name)
}

// Help is help message of version command
func (c *VersionCommand) Help() string {
	return ""
}
