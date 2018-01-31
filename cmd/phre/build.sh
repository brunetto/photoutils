#!/usr/bin/env bash

goimports -w *.go
rm -rf vendor
govendor init
#sed -i '' '/test/d' vendor/vendor.json && \
govendor add +external

go install

GOOS=darwin GOARCH=amd64 go build -o phre_darwin
GOOS=linux GOARCH=amd64 go build -o phre_linux
