#!/usr/bin/env bash

. /etc/profile

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o xshell2windterm.linux-arm64 xshell2windterm.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o xshell2windterm.linux-amd64 xshell2windterm.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o xshell2windterm.exe xshell2windterm.go
CGO_ENABLED=0 GOOS=linux GOARCH=darwin go build -o xshell2windterm.mac-amd64 xshell2windterm.go
