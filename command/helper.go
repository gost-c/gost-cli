package command

import (
	"github.com/mitchellh/go-homedir"
	"log"
	"path"
)

// BaseURL is the base url of gost api services
//var BaseURL = "http://localhost:8000/"
var BaseURL = "http://gost.congz.pw/"

// WebURL is the base url of gost web services
//var WebURL = "http://localhost:3000/"
var WebURL = "http://gost.surge.sh/#/"

const config = ".gostrc"

// ConfigFile is the true path of config file `.gostrc`
var ConfigFile string

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("An error occurred when get user home!")
	}
	ConfigFile = path.Join(home, config)
}
