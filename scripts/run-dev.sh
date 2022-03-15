#! /bin/sh

set -e


[ -x "$GOPATH/bin/gin" ]  ||  go install github.com/codegangsta/gin@latest


cd "$(dirname "$0")/../public"

export SUMBUR_CONFIG=sumbur-debug.yaml
gin  -a 8000  -b sumbur-debug  -t ../sumbur  run
