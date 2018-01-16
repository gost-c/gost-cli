package utils_test

import (
	"github.com/gost-c/gost-cli/utils"
	"testing"
)

func TestFileOrFolder(t *testing.T) {
	t.Log(utils.GetPathStat("../circle.yml"))
}
