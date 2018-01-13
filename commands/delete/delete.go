package delete

import (
	"github.com/gost-c/gost-cli/commands"
	"github.com/gost-c/gost-cli/utils"
	log "github.com/sirupsen/logrus"
	u "github.com/zcong1993/utils"
)

var url = utils.BaseURL + "api/delete/"

// Run is sub command runner for register
func Run(id string) {
	token, err := utils.ReadConfig()
	if err != nil || token == nil {
		log.Fatal("Get token failed, please login first")
	}
	url += id
	var res commands.Result
	err = u.GetJSONWithHeaders(url, &res, map[string]string{"Authorization": "Bearer " + string(token)})
	if err != nil {
		log.Fatalf("Unexpected error occurred: %s, make sure you have logged in", err.Error())
	}

	if res.Code != "200" {
		log.Fatalf("Delete error: %s", res.Msg)
	}

	log.Info(res.Msg)
}
