package delete

import (
	"fmt"
	"github.com/gost-c/gost-cli/commands"
	"github.com/gost-c/gost-cli/utils"
	u "github.com/zcong1993/utils"
)

var url = utils.BaseURL + "api/delete/"

// Run is sub command runner for register
func Run(id string) {
	token, err := utils.ReadConfig()
	if err != nil || token == nil {
		utils.Fail("Get token failed, please login first")
		return
	}
	url += id
	var res commands.Result
	err = u.GetJSONWithHeaders(url, &res, map[string]string{"Authorization": "Bearer " + string(token)})
	if err != nil {
		utils.Fail(fmt.Sprintf("Unexpected error occurred: %s, make sure you have logged in", err.Error()))
		return
	}

	if res.Code != "200" {
		utils.Fail(fmt.Sprintf("Delete error: %s", res.Msg))
		return
	}

	utils.Success(res.Msg)
}
