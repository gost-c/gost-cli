package register

import (
	"github.com/gost-c/gost-cli/commands"
	"github.com/gost-c/gost-cli/utils"
	log "github.com/sirupsen/logrus"
	u "github.com/zcong1993/utils"
)

var url = utils.BaseURL + "register"

// Run is sub command runner for register
func Run(username, password string) {
	user := commands.User{Username: username, Password: password}

	var res commands.Result

	err := u.PostJSON(url, user, &res, map[string]string{})

	if err != nil {
		log.Fatalf("Register error: %s", err.Error())
	}

	if res.Code != "200" {
		log.Fatalf("Register error: %s", res.Msg)
	}

	log.Info(res.Msg)
}
