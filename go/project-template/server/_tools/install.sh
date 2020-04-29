#!/usr/bin/env bash
set -e
cd _tools/tools
set -x
go run ../ls-imports/main.go -u -f main.go | xargs -tI % go install -v %