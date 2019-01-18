#!/bin/sh -e
APPNAME=pgbouncer-healthcheck
VERSIONFILE=version.go

<VERSION read -r version
build_date="$(date +%Y-%m-%d\ %H:%M)"

sed -i "$VERSIONFILE" -e 's/X\.X/'"$version"'/;s/DDMMYY/'"$build_date"'/'
CC=$(command -v gcc) go build -v --ldflags '-linkmode external -extldflags "-static -s"' -o "${APPNAME}"
