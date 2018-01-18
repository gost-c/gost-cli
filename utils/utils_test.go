package utils_test

import (
	"github.com/gost-c/gost-cli/utils"
	"testing"
)

func TestFileOrFolder(t *testing.T) {
	t.Log(utils.GetPathStat("../circle.yml"))
}

func TestLogPad(t *testing.T) {
	utils.LogPad("cdcdxsxs")
}

func TestGetToken(t *testing.T) {
	t.Log(utils.GetToken())
}
