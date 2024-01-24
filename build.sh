#!/bin/bash

# Build for Windows 64-bit
GOOS=windows GOARCH=amd64 go build -o build/goimer.exe cmd/cli/main.go

# Build for macOS 64-bit
GOOS=darwin GOARCH=amd64 go build -o build/goimer_darwin cmd/cli/main.go

# Build for Linux 64-bit
GOOS=linux GOARCH=amd64 go build -o build/goimer_linux cmd/cli/main.go
