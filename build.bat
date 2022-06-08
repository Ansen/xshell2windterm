@echo off

@REM linux arm
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=arm64
go build -o xshell2windterm.linux-arm64 xshell2windterm.go

@REM linux amd64
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o xshell2windterm.linux-amd64 xshell2windterm.go

@REM Win
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o xshell2windterm.exe xshell2windterm.go

@REM MAC
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o xshell2windterm.mac xshell2windterm.go