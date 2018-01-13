package utils

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"path"
)

var (
	// BaseURL is the base url of gost api services
	BaseURL = "http://gost.congz.pw/"
	// WebURL is the base url of gost web services
	WebURL = "http://gost.zcong.moe/#/gost/"
)

// ConfigFile is the true path of config file `.gostrc`
var ConfigFile string

const config = ".gostrc"

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("An error occurred when get user home!")
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
