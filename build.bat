@echo off
set GOARCH=wasm
set GOOS=js
set GOROOT_BOOTSTRAP=C:\Go
go build -o lib.wasm main.go