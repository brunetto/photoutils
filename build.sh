#!/usr/bin/env bash

cd cmd/phre 
go mod tidy
go install

GOOS=darwin GOARCH=amd64 go build -o phre_darwin
GOOS=linux GOARCH=amd64 go build -o phre_linux
