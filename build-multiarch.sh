#!/bin/bash

env GOOS=linux GOARCH=arm go build -o voe-dl-linux-arm src/*.go
env GOOS=linux GOARCH=amd64 go build -o voe-dl-linux-amd64 src/*.go
env GOOS=windows GOARCH=amd64 go build -o voe-dl-win-amd64 src/*.go
env GOOS=darwin GOARCH=amd64 go build -o voe-dl-macos-amd64 src/*.go

make