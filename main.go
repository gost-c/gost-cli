package main

import (
	"fmt"
	del "github.com/gost-c/gost-cli/commands/delete"
	"github.com/gost-c/gost-cli/commands/login"
	"github.com/gost-c/gost-cli/commands/push"
	"github.com/gost-c/gost-cli/commands/register"
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)

var (
	// Version is app version
	Version = "v0.1.0"
	// GitCommit is commit hash for version
	GitCommit = ""
)

var App = kingpin.New("gost", "Command line tool for gost.")

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

	deleteCmd = app.Command("delete", "Delete a gost you published.")
	id        = deleteCmd.Arg("id", "Gost id you want to delete.").Required().String()

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
	case deleteCmd.FullCommand():
		del.Run(*id)
	case version.FullCommand():
		showVersion()
	}
}

func showVersion() {
	version := fmt.Sprintf("\n%s version %s", app.Name, Version)
	if len(GitCommit) != 0 {
		version += fmt.Sprintf(" (%s)", GitCommit)
	}
	fmt.Println(version)
}
