package command

import (
	"strings"
	"github.com/zcong1993/utils"
	"io/ioutil"
)

// DeleteCommand is struct of delete meta
type DeleteCommand struct {
	Meta
}

// Run is entry function of delete command
func (c *DeleteCommand) Run(args []string) int {
	uflag := c.Meta.NewFlagSet("delete", c.Help())
	if err := uflag.Parse(args); err != nil {
		return 1
	}
	parsedArgs := uflag.Args()
	if len(parsedArgs) < 1 {
		c.UI.Error("Invalid argument: Usage gost delete HASH")
		return 1
	}
	hash := parsedArgs[0]
	url := BaseURL + "api/delete/" + hash
	token, err := ioutil.ReadFile(ConfigFile)
	if err != nil || string(token) == "" {
		c.UI.Error("Get token failed, please try login first")
		return 1
	}
	var res Result
	err = utils.GetJSONWithHeaders(url, &res, map[string]string{"Authorization": "Bearer " + string(token)})
	if err != nil {
		c.UI.Error("Unexpected error occurred, make sure you have logged in")
		return 1
	}
	if res.Code != "200" {
		c.UI.Error(res.Msg)
		return 1
	}
	c.UI.Info(res.Msg)
	return 0
}

// Synopsis is description of delete command
func (c *DeleteCommand) Synopsis() string {
	return "Delete a gost you published"
}

// Help is help message of delete command
func (c *DeleteCommand) Help() string {
	helpText := `
Delete a gost you have published.

Usage:

	gost delete HASH

`
	return strings.TrimSpace(helpText)
}
