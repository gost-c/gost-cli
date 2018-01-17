package main

import (
	"fmt"
	"github.com/gost-c/gost-cli/colors"
	del "github.com/gost-c/gost-cli/commands/delete"
	"github.com/gost-c/gost-cli/commands/folder"
	"github.com/gost-c/gost-cli/commands/login"
	"github.com/gost-c/gost-cli/commands/push"
	"github.com/gost-c/gost-cli/commands/register"
	"github.com/gost-c/gost-cli/commands/upgrade"
	"github.com/gost-c/gost-cli/utils"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	// GitCommit is commit hash for version
	GitCommit = ""
)

var (
	app = kingpin.New("gost", "Command line tool for gost.")

	registerCmd      = app.Command("register", "Register a account.")
	registerUsername = registerCmd.Flag("username", "Account username.").Short('u').Required().String()
	registerPassword = registerCmd.Flag("password", "Account password.").Short('p').Required().String()

	loginCmd      = app.Command("login", "Login your account.")
	loginUsername = loginCmd.Flag("username", "Account username.").Short('u').Required().String()
	loginPassword = loginCmd.Flag("password", "Account password.").Short('p').Required().String()

	pushCmd     = app.Command("push", "Push files to gost.")
	pushFiles   = pushCmd.Arg("files", "Push files.").Required().Strings()
	description = pushCmd.Flag("description", "Add some description").Short('d').Default("Published by zcong1993/gost.").String()

	folderCmd         = app.Command("folder", "Push folders to gost.")
	folderF           = folderCmd.Arg("folders", "Push folders.").Required().String()
	folderDescription = folderCmd.Flag("description", "Add some description").Short('d').Default("Published by zcong1993/gost.").String()
	notIgnoreHidden   = folderCmd.Flag("not-ignore-hidden", "Not ignore hidden folder").Short('n').Bool()

	deleteCmd = app.Command("delete", "Delete a gost you published.")
	id        = deleteCmd.Arg("id", "Gost id you want to delete.").Required().String()

	upgradeCmd = app.Command("upgrade", "Upgrade gost cli.")
	target     = upgradeCmd.Flag("target", "Upgrade to the specified version.").Short('t').String()

	version = app.Command("version", "Show gost cli version.")
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case registerCmd.FullCommand():
		register.Run(*registerUsername, *registerPassword)
	case loginCmd.FullCommand():
		login.Run(*loginUsername, *loginPassword)
	case pushCmd.FullCommand():
		push.Run(*pushFiles, *description)
	case folderCmd.FullCommand():
		folder.Run(*folderF, *folderDescription, *notIgnoreHidden)
	case deleteCmd.FullCommand():
		del.Run(*id)
	case version.FullCommand():
		showVersion()
	case upgradeCmd.FullCommand():
		upgrade.Run(*target)
	}
}

func showVersion() {
	version := fmt.Sprintf("%s version %s", colors.Cyan(app.Name), colors.Purple(utils.Version))
	if len(GitCommit) != 0 {
		version += colors.Gray(fmt.Sprintf(" (%s)", GitCommit))
	}
	utils.LogPad(version)
}
