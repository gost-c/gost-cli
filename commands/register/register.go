package register

import (
	"fmt"
	"github.com/gost-c/gost-cli/commands"
	"github.com/gost-c/gost-cli/utils"
	u "github.com/zcong1993/utils"
)

var url = utils.BaseURL + "api/register"

// Run is sub command runner for register
func Run(username, password string) {
	user := commands.User{Username: username, Password: password}

	var res commands.Result

	err := u.PostJSON(url, user, &res, map[string]string{})

	if err != nil {
		utils.Fail(fmt.Sprintf("Register error: %s", err.Error()))
		return
	}

	if !res.Success {
		utils.Fail(fmt.Sprintf("Register error: %s", res.Message))
		return
	}

	utils.Success(res.Data.(string))
}
