// Package upgrade forked from https://github.com/apex/up/blob/master/internal/cli/upgrade/upgrade.go
package upgrade

import (
	"fmt"
	"github.com/gost-c/gost-cli/colors"
	"github.com/gost-c/gost-cli/utils"
	"github.com/pkg/errors"
	"github.com/tj/go-update"
	"github.com/tj/go-update/progress"
	"github.com/tj/go-update/stores/github"
	"github.com/tj/go/term"
	"runtime"
	"time"
)

// Run is sub command runner for upgrade
func Run(target string) {
	doUpgrade(target)
}

func doUpgrade(target string) {
	v := normalizeVersion(utils.Version)

	start := time.Now()

	term.HideCursor()
	defer term.ShowCursor()

	p := &update.Manager{
		Command: "gost",
		Store: &github.Store{
			Owner:   "gost-c",
			Repo:    "gost-cli",
			Version: v,
		},
	}

	r, err := getLatestOrSpecified(p, target)

	if err != nil {
		utils.LogErrPad(errors.Wrap(err, "fetching release"))
		return
	}

	// no updates
	if r == nil {
		utils.LogSuccessPad("No updates available, you're good :)")
		return
	}

	a := r.FindTarball(runtime.GOOS, runtime.GOARCH)

	if a == nil {
		utils.LogErrPad(errors.Errorf("failed to find a binary for %s %s", runtime.GOOS, runtime.GOARCH))
		return
	}

	// download tarball to a tmp dir
	println()
	tarball, err := a.DownloadProxy(progress.Reader)
	if err != nil {
		utils.LogErrPad(errors.Wrap(err, "downloading tarball"))
		return
	}

	if err := p.Install(tarball); err != nil {
		utils.LogErrPad(errors.Wrap(err, "installing"))
		return
	}

	utils.LogSuccessPad(fmt.Sprintf("Updated %s to %s, cost %s", colors.Cyan(utils.Version), colors.Cyan(r.Version), colors.Yellow(time.Since(start).Round(time.Second).String())))
}

// getLatestOrSpecified returns the latest or specified release.
func getLatestOrSpecified(s update.Store, version string) (*update.Release, error) {
	if version == "" {
		return getLatest(s)
	}

	return s.GetRelease(version)
}

// getLatest returns the latest release, error, or nil when there is none.
func getLatest(s update.Store) (*update.Release, error) {
	releases, err := s.LatestReleases()

	if err != nil {
		return nil, errors.Wrap(err, "fetching releases")
	}

	if len(releases) == 0 {
		return nil, nil
	}

	return releases[0], nil
}

func normalizeVersion(version string) string {
	if len(version) < 1 {
		return version
	}
	if version[:1] == "v" {
		return version[1:]
	}
	return version
}
