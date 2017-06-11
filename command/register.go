package command

import (
	"github.com/zcong1993/utils"
	"strings"
)

// RegisterCommand is struct of register meta
type RegisterCommand struct {
	Meta
}

// User is struct of login api's post data
type User struct {
	Username string
	Password string
}

// Result is struct of common api's response
type Result struct {
	Code string `decoder:"code"`
	Msg  string `decoder:"msg"`
}

// Run is entry function of register command
func (c *RegisterCommand) Run(args []string) int {
	var username, password string
	uflag := c.Meta.NewFlagSet("register", c.Help())

	uflag.StringVar(&username, "username", "", "username")
	uflag.StringVar(&username, "u", "", "username (short)")

	uflag.StringVar(&password, "password", "", "password")
	uflag.StringVar(&password, "p", "", "password (short)")

	if err := uflag.Parse(args); err != nil {
		return 1
	}
	if username == "" || password == "" {
		c.UI.Error("Invalid options: Usage gost register -u=yourUsername -p=yourPassword")
		return 1
	}
	url := BaseURL + "register"
	user := User{username, password}
	var res Result
	err := utils.PostJSON(url, user, &res, map[string]string{})
	if err != nil {
		c.UI.Error("Unexpected error occurred, please try again")
		return 1
	}
	if res.Code != "200" {
		c.UI.Error(res.Msg)
		return 1
	}
	c.UI.Info(res.Msg)
	return 0
}

// Synopsis is description of register command
func (c *RegisterCommand) Synopsis() string {
	return "Register a account"
}

// Help is help message of register command
func (c *RegisterCommand) Help() string {
	helpText := `
Before using gost, you should register a account first.

Usage:

	gost register [options]

Options:

	-username=yourname, -u          Choose a username, length should > 6 and < 20
	-password=yourpassword, -p      Choose a password, length should > 6

Example:
	gost register -u=gostuser1 -p=youneverknow
`
	return strings.TrimSpace(helpText)
}
