package utils

import (
	"fmt"
	"github.com/gost-c/gost-cli/colors"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path"
)

var (
	// BaseURL is the base url of gost api services
	BaseURL = "http://gost.congz.pw/"
	// WebURL is the base url of gost web services
	WebURL = "http://gost.zcong.moe/#/gost/"
	space  = "  "
)

// ConfigFile is the true path of config file `.gostrc`
var ConfigFile string

const config = ".gostrc"

func init() {
	home, err := homedir.Dir()
	if err != nil {
		Fail("An error occurred when get user home!")
		os.Exit(1)
	}
	ConfigFile = path.Join(home, config)
}

// WriteConfig write token to config file
func WriteConfig(token []byte) error {
	return ioutil.WriteFile(ConfigFile, token, 0644)
}

// ReadConfig read token from config file
func ReadConfig() ([]byte, error) {
	return ioutil.ReadFile(ConfigFile)
}

// Success log success message with colors
func Success(str string) {
	fmt.Println()
	fmt.Printf("%s%s%s", colors.Green("SUCCESS"), space, str)
	fmt.Println()
}

// Fail log error message with colors
func Fail(str string) {
	fmt.Println()
	fmt.Printf("%s%s%s", colors.Red("ERROR"), space, str)
	fmt.Println()
}
