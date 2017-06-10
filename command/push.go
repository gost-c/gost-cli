package command

import (
	"strings"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"fmt"
	"os"
	"path"
	"github.com/pkg/errors"
	"github.com/zcong1993/utils"
)

type PushCommand struct {
	Meta
}

type Gist struct {
	Public bool
	Description string
	Version uint
	Files []File
}

type File struct {
	Filename string
	Content string
}

func getFiles(files []string) ([]File, error) {
	eg := errgroup.Group{}
	var results []File
	for _, file := range files {
		file := file
		eg.Go(func() error {
			content, err := ioutil.ReadFile(file)
			fmt.Fprintf(os.Stdout, "--> Parsing file: %15s\n", file)
			if err != nil {
				return errors.Wrapf(err,
					"failed to get file content: %s", file)
			}
			results = append(results, File{path.Base(file), string(content)})
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, errors.Wrap(err, "one of the goroutines failed")
	}
	return results, nil
}

func (c *PushCommand) Run(args []string) int {
	var public bool
	var description string
	uflag := c.Meta.NewFlagSet("push", c.Help())

	uflag.BoolVar(&public, "public", false, "Make the gost public")
	uflag.BoolVar(&public, "p", false, "Make the gost public (short)")

	uflag.StringVar(&description, "description", "Pushed by zcong1993/gost.", "Add some description")
	uflag.StringVar(&description, "d", "Pushed by zcong1993/gost.", "Add some description (short)")

	if err := uflag.Parse(args); err != nil {
		return 1
	}

	parsedArgs := uflag.Args()
	if len(parsedArgs) < 1 {
		c.Ui.Error("Invalid argument: Usage glic push [options] FILE1 FILE2 ...")
		return 1
	}
	files, err := getFiles(parsedArgs)
	if err != nil {
		c.Ui.Error(err.Error())
		return 1
	}
	gist := Gist{public, description, 1, files}
	url := BaseUrl + "api/create"
	token, err := ioutil.ReadFile(ConfigFile)
	if err != nil || string(token) == "" {
		c.Ui.Error("Get token failed, please try login first")
		return 1
	}
	var res Result
	err = utils.PostJSON(url, gist, &res, map[string]string{"Authorization": "Bearer " + string(token)})
	if err != nil {
		c.Ui.Error("Unexpected error occurred, make sure you have logged in")
		return 1
	}
	if res.Code != "200" {
		c.Ui.Error(res.Msg)
		return 1
	}
	c.Ui.Info("Push success! The url is " + WebUrl + res.Msg)
	return 0
}

func (c *PushCommand) Synopsis() string {
	return "Push some doc to gost"
}

func (c *PushCommand) Help() string {
	helpText := `
You can use gost push some doc online to share with friends quickly.

Usage:

	gost push [options] FILE1 FILE2 ...

Options:

	-public, -p                                 Make the gost public (not support yet)
	-description=description, -d=description    Add some description, default is "Pushed by zcong1993/gost"

Example:

	gost main.go README.md
`
	return strings.TrimSpace(helpText)
}
