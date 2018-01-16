package folder

import (
	"fmt"
	"github.com/gost-c/gost-cli/colors"
	"github.com/gost-c/gost-cli/commands/push"
	"github.com/gost-c/gost-cli/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Run is sub command runner for push
func Run(path string, description string, notIgnoreFolder bool) {
	f, err := GetFiles(path, notIgnoreFolder)
	if err != nil {
		utils.Fail(fmt.Sprintf("Parse files error: %s", colors.Red(err.Error())))
		return
	}
	push.PushGost(f, description)
}

// GetFiles get files from path
func GetFiles(path string, notIgnoreFolder bool) ([]push.File, error) {
	var fs []push.File

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if isHiddenPath(path) && !notIgnoreFolder {
				fmt.Printf("%s Skip hidden folder: %15s\n", colors.Blue("-->"), colors.Red(path))
				return filepath.SkipDir
			}
			fmt.Printf("%s Parse folder: %15s\n", colors.Blue("-->"), colors.Purple(path))
			return nil
		} else {
			if len(fs) > utils.MaxFilesCount {
				return utils.ErrMaxFilesCount
			}
			fmt.Printf("%s Reading file: %15s\n", colors.Blue("-->"), colors.Cyan(path))
			content, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}
			f := push.File{
				Filename: path,
				Content:  string(content),
			}
			fs = append(fs, f)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func isHiddenPath(p string) bool {
	arr := strings.Split(p, "/")
	for _, v := range arr {
		if v == "." || v == ".." || v == "" {
			continue
		}
		if v[:1] == "." {
			return true
		}
	}
	return false
}
