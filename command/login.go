package command

import (
	"github.com/zcong1993/utils"
	"io/ioutil"
	"strings"
)

// LoginCommand is struct of login meta
type LoginCommand struct {
	Meta
}

// LoginResult is struct of login api's response
type LoginResult struct {
	Expire string `decoder:"expire"`
	Token  string `decoder:"token"`
}

// Run is entry function of login command
func (c *LoginCommand) Run(args []string) int {
	var username, password string
	uflag := c.Meta.NewFlagSet("login", c.Help())

	uflag.StringVar(&username, "username", "", "username")
	uflag.StringVar(&username, "u", "", "username (short)")

	uflag.StringVar(&password, "password", "", "password")
	uflag.StringVar(&password, "p", "", "password (short)")

	if err := uflag.Parse(args); err != nil {
		return 1
	}
	if username == "" || password == "" {
		c.UI.Error("Invalid options: Usage gost login -u=yourUsername -p=yourPassword")
		return 1
	}

	url := BaseURL + "login"
	user := User{username, password}
	var res LoginResult
	err := utils.PostJSON(url, user, &res, map[string]string{})
	if err != nil {
		c.UI.Error("Unexpected error occurred, please try again")
		return 1
	}
	if res.Token == "" {
		c.UI.Error("Username or password is wrong, please try again")
		return 1
	}
	err = ioutil.WriteFile(ConfigFile, []byte(res.Token), 0644)
	if err != nil {
		c.UI.Error("Unexpected error occurred when write config file")
		return 1
	}
	c.UI.Info("Success! Now you can use `gost push`")
	return 0
}

// Synopsis is description of login command
func (c *LoginCommand) Synopsis() string {
	return "Login your account"
}

// Help is help message of login command
func (c *LoginCommand) Help() string {
	helpText := `
Using gost need token, you can get it via login.

Usage:

	gost login [options]

Options:

	-username=yourname, -u          Type your username
	-password=yourpassword, -p      Type your password

Example:
	gost login -u=gostuser1 -p=youneverknow
`
	return strings.TrimSpace(helpText)
}
