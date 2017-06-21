package command

import (
	"bytes"
	"fmt"
	latest "github.com/tcnksm/go-latest"
	"time"
)

const defaultCheckTimeout = 2 * time.Second

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
	var buf bytes.Buffer
	verCheckCh := make(chan *latest.CheckResponse)
	go func() {
		fixFunc := latest.DeleteFrontV()
		githubTag := &latest.GithubTag{
			Owner:             "gost-c",
			Repository:        "gost-cli",
			FixVersionStrFunc: fixFunc,
		}

		res, err := latest.Check(githubTag, fixFunc(c.Version))
		if err != nil {
			// Don't return error
			return
		}
		verCheckCh <- res
	}()

	select {
	case <-time.After(defaultCheckTimeout):
	case res := <-verCheckCh:
		if res.Outdated {
			fmt.Fprintf(&buf,
				"Latest version of %s is v%s, please upgrade!\n",
				c.Name, res.Current)
		}
	}
	c.UI.Output(buf.String())
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
