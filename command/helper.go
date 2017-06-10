package command

import (
	"github.com/mitchellh/go-homedir"
	"log"
	"path"
)

var BaseUrl = "http://localhost:8000/"
var WebUrl = "http://localhost:3000/"

const config = ".gostrc"
var ConfigFile string

func init() {
	home, err := homedir.Dir()
	if err != nil {
		log.Fatal("An error occured when get user home!")
	}
	ConfigFile = path.Join(home, config)
}