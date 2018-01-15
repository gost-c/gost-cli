package push

import (
	"fmt"
	"github.com/gost-c/gost-cli/colors"
	"github.com/gost-c/gost-cli/commands"
	"github.com/gost-c/gost-cli/utils"
	"github.com/pkg/errors"
	u "github.com/zcong1993/utils"
	"io/ioutil"
	"path"
)

var url = utils.BaseURL + "api/gost"

// Gost is struct of push api's post data
type Gost struct {
	Public      bool
	Description string
	Files       []File
}

// File is struct of gist's file
type File struct {
	Filename string
	Content  string
}

// Run is sub command runner for push
func Run(files []string, description string) {
	f, err := getFiles(files)
	if err != nil {
		utils.Fail(fmt.Sprintf("parse files error: %v", err))
		return
	}
	gist := Gost{Public: true, Description: description, Files: f}
	token, err := utils.ReadConfig()
	if err != nil || token == nil {
		utils.Fail("Get token failed, please login first")
		return
	}

	var res commands.Result
	err = u.PostJSON(url, gist, &res, map[string]string{"Authorization": "Bearer " + string(token)})
	if err != nil {
		utils.Fail(fmt.Sprintf("Unexpected error occurred: %s, make sure you have logged in", err.Error()))
		return
	}

	if !res.Success {
		utils.Fail(fmt.Sprintf("Push error: %s", res.Message))
		return
	}
	utils.Success(fmt.Sprintf("Push success! The url is %s", colors.Yellow(utils.WebURL+res.Data.(string))))
}

func getFiles(files []string) ([]File, error) {
	var results []File
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		fmt.Printf("%s Reading file: %15s\n", colors.Blue("-->"), colors.Cyan(file))
		if err != nil {
			return nil, errors.Wrapf(err,
				"failed to get file content: %s", file)
		}
		results = append(results, File{path.Base(file), string(content)})
	}
	return results, nil
}
