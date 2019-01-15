#!/bin/sh -e

go get github.com/golang/dep/cmd/dep
dep ensure -v --vendor-only
