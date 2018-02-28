package folder_test

import (
	"testing"

	"github.com/gost-c/gost-cli/commands/folder"
)

func TestGetFiles(t *testing.T) {
	fs, err := folder.GetFiles("../", false, []string{})
	t.Log(len(fs), err)
}
