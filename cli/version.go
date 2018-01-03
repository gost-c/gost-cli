package cli

// Name is cli's name
const Name string = "gost-cli"

// Version is cli's current version
const Version string = "v1.2.0"

// GitCommit describes latest commit hash.
// This value is extracted by git command when building.
// To set this from outside, use go build -ldflags "-X main.GitCommit \"$(COMMIT)\""
var GitCommit string
