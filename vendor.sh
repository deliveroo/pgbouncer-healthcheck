#!/bin/sh -e

go get github.com/golang/dep/cmd/dep
(cd ..; dep ensure -v --vendor-only)
