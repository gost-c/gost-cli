package login

import (
	"github.com/gost-c/gost-cli/commands"
	"github.com/gost-c/gost-cli/utils"
	log "github.com/sirupsen/logrus"
	u "github.com/zcong1993/utils"
)

// LoginResult is struct of login api's response
type LoginResult struct {
	Expire string `decoder:"expire"`
	Token  string `decoder:"token"`
}

var url = utils.BaseURL + "login"

// Run is sub command runner for login
func Run(username, password string) {

	user := commands.User{Username: username, Password: password}

	var res LoginResult
	err := u.PostJSON(url, user, &res, map[string]string{})

	if err != nil {
		log.Fatalf("Login error: %s, please try again.", err.Error())
	}

	if res.Token == "" {
		log.Fatal("Username or password is wrong, please try again")
	}

	err = utils.WriteConfig([]byte(res.Token))

	if err != nil {
		log.Fatalf("Unexpected error occurred when write config file: %s", err.Error())
	}

	log.Info("Success! Now you can use `gost push`")
}
