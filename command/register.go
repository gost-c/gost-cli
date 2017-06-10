package command

import (
	"github.com/zcong1993/utils"
	"strings"
)

type RegisterCommand struct {
	Meta
}

type User struct {
	Username string
	Password string
}

type Result struct {
	Code string `decoder:"code"`
	Msg  string `decoder:"msg"`
}

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
		c.Ui.Error("Invalid options: Usage gost register -u=yourUsername -p=yourPassword")
		return 1
	}
	url := BaseUrl + "register"
	user := User{username, password}
	var res Result
	err := utils.PostJSON(url, user, &res, map[string]string{})
	if err != nil {
		c.Ui.Error("Unexpected error occurred, please try again")
		return 1
	}
	if res.Code != "200" {
		c.Ui.Error(res.Msg)
		return 1
	}
	c.Ui.Info(res.Msg)
	return 0
}

func (c *RegisterCommand) Synopsis() string {
	return "Register a account"
}

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
