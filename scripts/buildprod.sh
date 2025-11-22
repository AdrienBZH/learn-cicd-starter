#!/bin/bash
set -e

export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

go env GOOS
go env GOARCH

go build -o notely