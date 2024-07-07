#!/usr/local/bin/bash

GOARCH=amd64 GOOS=openbsd govvv build -ldflags="-s -w" -o binaries/ago-openbsd-amd64 .
GOARCH=386 GOOS=openbsd govvv build -ldflags="-s -w" -o binaries/ago-openbsd-i386 .
GOARCH=arm GOOS=openbsd govvv build -ldflags="-s -w" -o binaries/ago-openbsd-arm .
GOARCH=arm64 GOOS=openbsd govvv build -ldflags="-s -w" -o binaries/ago-openbsd-arm64 .

GOARCH=amd64 GOOS=linux govvv build -ldflags="-s -w" -o binaries/ago-linux-amd64 .
GOARCH=386 GOOS=linux govvv build -ldflags="-s -w" -o binaries/ago-linux-i386 .
GOARCH=arm GOOS=linux govvv build -ldflags="-s -w" -o binaries/ago-linux-arm .
GOARCH=arm64 GOOS=linux govvv build -ldflags="-s -w" -o binaries/ago-linux-arm64 .
