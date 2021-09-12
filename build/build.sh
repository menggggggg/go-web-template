#!/bin/sh

set -xe

git config --global url."git@github.com:".insteadOf "https://github.com"

mkdir -p /dist

export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
# export GOPRIVATE=

go mod download 

APPNAME=go-web-template
VERSION=$(git describe --tags)
BUILDTIME=`date + "+%Y%m%d%I%M%S"`
GITSHA1=$(git rev-parse HEAD)

REPO="github.com/menggggggg/go-web-template"

GO_LDFLAGS="-X ${REPO}/pkg/util.Name=${APPNAME} \
    -X ${REPO}/pkg/Version=${VERSION} \
    -X ${REPO}/pkg/BuildTime=${BUILDTIME} \
    -X ${REPO}/pkg/GitSHA1=${GITSHA1}"


CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "${GO_LDFLAGS}" -o dist/&{APPNAME} main.go