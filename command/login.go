package command

import (
	"github.com/zcong1993/utils"
	"io/ioutil"
	"strings"
)

type LoginCommand struct {
	Meta
}

type LoginResult struct {
	Expire string `decoder:"expire"`
	Token  string `decoder:"token"`
}

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
		c.Ui.Error("Invalid options: Usage gost login -u=yourUsername -p=yourPassword")
		return 1
	}

	url := BaseUrl + "login"
	user := User{username, password}
	var res LoginResult
	err := utils.PostJSON(url, user, &res, map[string]string{})
	if err != nil {
		c.Ui.Error("Unexpected error occurred, please try again")
		return 1
	}
	if res.Token == "" {
		c.Ui.Error("Username or password is wrong, please try again")
		return 1
	}
	err = ioutil.WriteFile(ConfigFile, []byte(res.Token), 0644)
	if err != nil {
		c.Ui.Error("Unexpected error occurred when write config file")
		return 1
	}
	c.Ui.Info("Success! Now you can use `gost push`")
	return 0
}

func (c *LoginCommand) Synopsis() string {
	return "Login your account"
}

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
