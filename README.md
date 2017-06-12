# gost-cli

[![Go Report Card](https://goreportcard.com/badge/github.com/gost-c/gost-cli)](https://goreportcard.com/report/github.com/gost-c/gost-cli)
[![Build Status](https://travis-ci.org/gost-c/gost-cli.svg?branch=master)](https://travis-ci.org/gost-c/gost-cli)

> gost-cli is a command line tool for gost

## Description

`gost` is a gist-like services (see the sample [sample link](http://gost.surge.sh/#/7f6fbcc7-8a8b-443e-a88e-39f49c693215)), and `gost-cli` is a command tool for it.

## Usage

```bash
# show help
$ gost -h
# show sub command help, example gost login -h
$ gost <sub command> -h
```

## Install

### Download directly (recommend)

[Download](https://github.com/gost-c/gost-cli/releases) the package (should match your platform) and move it to any `$PATH` folder.

### To install, use `go get`:

```bash
$ go get -d github.com/gost-c/gost-cli
```

## Docs

see [http://gost-docs.congz.pw](http://gost-docs.congz.pw)

## Contribution

1. Fork ([https://github.com/gost-c/gost-cli/fork](https://github.com/gost-c/gost-cli/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run test suite with the `go test ./...` command and confirm that it passes
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[zcong1993](https://github.com/zcong1993)
