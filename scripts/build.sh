#!/usr/bin/env bash

GOOS=darwin GOARCH=arm64 go build -o bin/spacer-darwin-arm64 .
GOOS=linux GOARCH=amd64 go build -o bin/spacer-linux-amd64 .
