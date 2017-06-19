#!/bin/bash
set -ev

XC_ARCH=${XC_ARCH:-amd64}
XC_OS=${XC_OS:-darwin linux windows}

COMMIT=`git describe --always`


rm -rf pkg/
gox \
    -ldflags "-X main.GitCommit=${COMMIT}" \
    -parallel=-1 \
    -os="${XC_OS}" \
    -arch="${XC_ARCH}" \
    -output "pkg/gost-{{.OS}}"