package login

import (
	"fmt"
	"github.com/gost-c/gost-cli/commands"
	"github.com/gost-c/gost-cli/utils"
	u "github.com/zcong1993/utils"
)

var url = utils.BaseURL + "api/login"

// Run is sub command runner for login
func Run(username, password string) {

	user := commands.User{Username: username, Password: password}

	var res commands.Result
	err := u.PostJSON(url, user, &res, map[string]string{})

	if err != nil {
		utils.Fail(fmt.Sprintf("Login error: %s, please try again.", err.Error()))
		return
	}

	if !res.Success {
		utils.Fail(fmt.Sprintf("Login error: %s, please try again.", res.Message))
		return
	}

	if res.Data == nil {
		utils.Fail("Username or password is wrong, please try again")
		return
	}

	err = utils.WriteConfig([]byte(res.Data.(string)))

	if err != nil {
		utils.Fail(fmt.Sprintf("Unexpected error occurred when write config file: %s", err.Error()))
		return
	}

	utils.Success("Login success. Now you can use `gost push`")
}
