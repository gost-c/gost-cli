package command

import (
	"bytes"
	"fmt"
)

// VersionCommand is struct of version meta
type VersionCommand struct {
	Meta

	Name     string
	Version  string
	Revision string
}

// Run is entry function of version command
func (c *VersionCommand) Run(args []string) int {
	var versionString bytes.Buffer

	fmt.Fprintf(&versionString, "%s version %s", c.Name, c.Version)
	if c.Revision != "" {
		fmt.Fprintf(&versionString, " (%s)", c.Revision)
	}

	c.UI.Output(versionString.String())
	return 0
}

// Synopsis is description of version command
func (c *VersionCommand) Synopsis() string {
	return fmt.Sprintf("Print %s version and quit", c.Name)
}

// Help is help message of version command
func (c *VersionCommand) Help() string {
	return ""
}
