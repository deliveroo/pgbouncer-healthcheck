#!/bin/sh -e
APPNAME=pgbouncer-healthcheck

CC=$(command -v gcc) go build -v --ldflags '-linkmode external -extldflags "-static -s"' -o "${APPNAME}"
