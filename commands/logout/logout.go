package logout

import "github.com/gost-c/gost-cli/utils"

// Run is sub command runner for logout
func Run() {
	err := utils.WriteConfig([]byte(""))
	if err != nil {
		utils.LogErrPad(err)
	}
	utils.LogSuccessPad("Now you logout. ")
}
