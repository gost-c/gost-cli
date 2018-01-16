package folder_test

import (
	"github.com/gost-c/gost-cli/commands/folder"
	"testing"
)

func TestGetFiles(t *testing.T) {
	fs, err := folder.GetFiles("../")
	t.Log(len(fs), err)
}
