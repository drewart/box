#!/bin/bash
#
source ./version.env

if [ -z "$VERSION" ]; then
  echo "VERSION not set in"
  exit 1
fi

go build -o bin/box -ldflags "-X main.version=${VERSION}" -trimpath -v ./cmd/box
