#!/bin/bash

GOARCH=amd64 GOOS=openbsd govvv build -ldflags="-s -w" -o binaries/ago-openbsd-amd64 ago.go
GOARCH=386 GOOS=openbsd govvv build -ldflags="-s -w" -o binaries/ago-openbsd-i386 ago.go
GOARCH=arm GOOS=openbsd govvv build -ldflags="-s -w" -o binaries/ago-openbsd-arm ago.go
GOARCH=arm64 GOOS=openbsd govvv build -ldflags="-s -w" -o binaries/ago-openbsd-arm64 ago.go

GOARCH=amd64 GOOS=linux govvv build -ldflags="-s -w" -o binaries/ago-linux-amd64 ago.go
GOARCH=386 GOOS=linux govvv build -ldflags="-s -w" -o binaries/ago-linux-i386 ago.go
GOARCH=arm GOOS=linux govvv build -ldflags="-s -w" -o binaries/ago-linux-arm ago.go
GOARCH=arm64 GOOS=linux govvv build -ldflags="-s -w" -o binaries/ago-linux-arm64 ago.go
