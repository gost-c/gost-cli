package url

import (
	"fmt"
	"github.com/gost-c/gost-cli/colors"
	"github.com/gost-c/gost-cli/utils"
	"github.com/pkg/browser"
)

// Run is sub command runner for url
func Run(open bool) {
	utils.LogPad(fmt.Sprintf("%s  %s", colors.Blue("URL"), colors.Yellow(utils.WebURL)))
	if open {
		browser.OpenURL(utils.WebURL)
	}
}
